package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client is the client for the ACS API.
type Client struct {
	base *resty.Client
	key  string
}

// Opt is a type for options.
type Opt func(*resty.Request)

// WithAuthToken sets the auth token.
func WithAuthToken(token string) Opt {
	return func(r *resty.Request) {
		r.SetAuthToken(token)
	}
}

// New creates a new Client.
func New(endpointURL, key string, c *http.Client) *Client {
	base := resty.
		NewWithClient(c).
		SetBaseURL(endpointURL)

	return &Client{base, key}
}

// R ...
func (c *Client) R() *resty.Request {
	return c.base.R()
}

// Post ...
func (c *Client) Post(ctx context.Context, key string, resource string, query string, reqbody any, response any, opts ...Opt) error {
	body := []byte("{}")

	var err error

	if reqbody != nil {
		body, err = json.Marshal(reqbody)
		if err != nil {
			return err
		}
	} else {
		reqbody = struct{}{}
	}

	u, err := url.Parse(c.base.BaseURL)
	if err != nil {
		return err
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	contentHash := computeContentHash(body)
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", resource+"?"+query, date, u.Host, contentHash)

	signature := computeSignature(stringToSign, key)

	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)

	res, err := c.base.R().
		EnableTrace().
		SetContext(ctx).
		SetHeader("x-ms-date", date).
		SetHeader("x-ms-content-sha256", contentHash).
		SetHeader("Authorization", authorizationHeader).
		SetHeader("Content-Type", "application/json").
		SetBody(reqbody).
		SetResult(response).
		SetQueryString(query).
		Post(resource)
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		fmt.Println(res)
		return fmt.Errorf("failed to send request: %s", res.String())
	}

	return nil
}

// Delete ...
func (c *Client) Delete(ctx context.Context, resource string, query string, response any) error {
	body := []byte("{}")
	reqBody := struct{}{}

	date := time.Now().UTC().Format(http.TimeFormat)

	u, err := url.Parse(c.base.BaseURL)
	if err != nil {
		return err
	}

	contentHash, authHeader := createAuthHeader(
		"DELETE",
		u.Host,
		resource+"?"+query,
		date,
		c.key,
		body,
	)

	res, err := c.base.R().
		EnableTrace().
		SetContext(ctx).
		SetBody(reqBody).
		SetHeader("x-ms-date", date).
		SetHeader("x-ms-content-sha256", contentHash).
		SetHeader("Authorization", authHeader).
		SetHeader("Content-Type", "application/json").
		SetResult(response).
		SetQueryString(query).
		Delete(resource)
	if err != nil {
		return err
	}

	if !res.IsSuccess() {
		return fmt.Errorf("failed to send request: %s", res.String())
	}

	return nil
}

func createAuthHeader(method string, host string, resourcePath string, date string, secret string, body []byte) (string, string) {
	contentHash := computeContentHash(body)
	stringToSign := fmt.Sprintf("%s\n%s\n%s;%s;%s", method, resourcePath, date, host, contentHash)
	signature := computeSignature(stringToSign, secret)

	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)

	return contentHash, authorizationHeader
}

func computeContentHash(content []byte) string {
	sha256 := sha256.New()

	sha256.Write(content)
	hashedBytes := sha256.Sum(nil)
	base64EncodedBytes := base64.StdEncoding.EncodeToString(hashedBytes)

	return base64EncodedBytes
}

func computeSignature(stringToSign string, secret string) string {
	decodedSecret, _ := base64.StdEncoding.DecodeString(secret)
	hash := hmac.New(sha256.New, decodedSecret)

	hash.Write([]byte(stringToSign))
	hashedBytes := hash.Sum(nil)
	encodedSignature := base64.StdEncoding.EncodeToString(hashedBytes)

	return encodedSignature
}

package client

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"

	goquery "github.com/google/go-querystring/query"
)

const (
	contentType        = "Content-Type"
	jsonContentType    = "application/json"
	formContentType    = "application/x-www-form-urlencoded"
	signedHeaderPrefix = "HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature="
)

// DefaultVersion
var DefaultVersion = struct {
	APIVersion string `url:"api-version"`
}{
	APIVersion: "2024-06-15-preview",
}

// Doer executes http requests.  It is implemented by *http.Client.  You can
// wrap *http.Client with layers of Doers to form a stack of client-side
// middleware.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is an HTTP Request builder and sender.
type Client struct {
	// http Client for doing requests
	httpClient Doer
	// HTTP method (GET, POST, etc.)
	method string
	// raw url string for requests
	rawURL string
	// stores key-values pairs to add to request's Headers
	header http.Header
	// url tagged query structs
	queryStructs []interface{}
	// body provider
	bodyProvider BodyProvider
	// response decoder
	responseDecoder ResponseDecoder
	// signer
	signer SignerProvider
}

// New returns a new Client with an http DefaultClient.
func New() *Client {
	return &Client{
		httpClient:      http.DefaultClient,
		method:          "GET",
		header:          make(http.Header),
		queryStructs:    make([]interface{}, 0),
		responseDecoder: jsonDecoder{},
		signer:          noopSigner{},
	}
}

// New returns a copy of a Client for creating a new Client with properties
// from a parent Client. For example,
//
//	parentClient := Client.New().Client(client).Base("https://api.io/")
//	fooClient := parentClient.New().Get("foo/")
//	barClient := parentClient.New().Get("bar/")
//
// fooClient and barClient will both use the same client, but send requests to
// https://api.io/foo/ and https://api.io/bar/ respectively.
//
// Note that query and body values are copied so if pointer values are used,
// mutating the original value will mutate the value within the child Client.
func (s *Client) New() *Client {
	// copy Headers pairs into new Header map
	headerCopy := make(http.Header)
	for k, v := range s.header {
		headerCopy[k] = v
	}

	return &Client{
		httpClient:      s.httpClient,
		method:          s.method,
		rawURL:          s.rawURL,
		header:          headerCopy,
		queryStructs:    append([]interface{}{}, s.queryStructs...),
		bodyProvider:    s.bodyProvider,
		responseDecoder: s.responseDecoder,
		signer:          s.signer,
	}
}

// Http Client

// Client sets the http Client used to do requests. If a nil client is given,
// the http.DefaultClient will be used.
func (s *Client) Client(httpClient *http.Client) *Client {
	if httpClient == nil {
		return s.Doer(http.DefaultClient)
	}
	return s.Doer(httpClient)
}

// Doer sets the custom Doer implementation used to do requests.
// If a nil client is given, the http.DefaultClient will be used.
func (s *Client) Doer(doer Doer) *Client {
	if doer == nil {
		s.httpClient = http.DefaultClient
	} else {
		s.httpClient = doer
	}
	return s
}

// Method

// Head sets the Client method to HEAD and sets the given pathURL.
func (s *Client) Head(pathURL string) *Client {
	s.method = "HEAD"
	return s.Path(pathURL)
}

// Get sets the Client method to GET and sets the given pathURL.
func (s *Client) Get(pathURL string) *Client {
	s.method = "GET"
	return s.Path(pathURL)
}

// Post sets the Client method to POST and sets the given pathURL.
func (s *Client) Post(pathURL string) *Client {
	s.method = "POST"
	return s.Path(pathURL)
}

// Put sets the Client method to PUT and sets the given pathURL.
func (s *Client) Put(pathURL string) *Client {
	s.method = "PUT"
	return s.Path(pathURL)
}

// Patch sets the Client method to PATCH and sets the given pathURL.
func (s *Client) Patch(pathURL string) *Client {
	s.method = "PATCH"
	return s.Path(pathURL)
}

// Delete sets the Client method to DELETE and sets the given pathURL.
func (s *Client) Delete(pathURL string) *Client {
	s.method = "DELETE"
	return s.Path(pathURL)
}

// Options sets the Client method to OPTIONS and sets the given pathURL.
func (s *Client) Options(pathURL string) *Client {
	s.method = "OPTIONS"
	return s.Path(pathURL)
}

// Trace sets the Client method to TRACE and sets the given pathURL.
func (s *Client) Trace(pathURL string) *Client {
	s.method = "TRACE"
	return s.Path(pathURL)
}

// Connect sets the Client method to CONNECT and sets the given pathURL.
func (s *Client) Connect(pathURL string) *Client {
	s.method = "CONNECT"
	return s.Path(pathURL)
}

// Header

// Add adds the key, value pair in Headers, appending values for existing keys
// to the key's values. Header keys are canonicalized.
func (s *Client) Add(key, value string) *Client {
	s.header.Add(key, value)
	return s
}

// Set sets the key, value pair in Headers, replacing existing values
// associated with key. Header keys are canonicalized.
func (s *Client) Set(key, value string) *Client {
	s.header.Set(key, value)
	return s
}

// SignProvider sets the Client's SignerProvider.
func (s *Client) SignProvider(signer SignerProvider) *Client {
	if signer == nil {
		return s
	}

	s.signer = signer

	return s
}

// SetBasicAuth sets the Authorization header to use HTTP Basic Authentication
// with the provided username and password. With HTTP Basic Authentication
// the provided username and password are not encrypted.
func (s *Client) SetBasicAuth(username, password string) *Client {
	return s.Set("Authorization", "Basic "+basicAuth(username, password))
}

// basicAuth returns the base64 encoded username:password for basic auth copied
// from net/http.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// Url

// Base sets the rawURL. If you intend to extend the url with Path,
// baseUrl should be specified with a trailing slash.
func (s *Client) Base(rawURL string) *Client {
	s.rawURL = rawURL
	return s
}

// Path extends the rawURL with the given path by resolving the reference to
// an absolute URL. If parsing errors occur, the rawURL is left unmodified.
func (s *Client) Path(path string) *Client {
	baseURL, baseErr := url.Parse(s.rawURL)
	pathURL, pathErr := url.Parse(path)
	if baseErr == nil && pathErr == nil {
		s.rawURL = baseURL.ResolveReference(pathURL).String()
		return s
	}
	return s
}

// QueryStruct appends the queryStruct to the Client's queryStructs. The value
// pointed to by each queryStruct will be encoded as url query parameters on
// new requests (see Request()).
// The queryStruct argument should be a pointer to a url tagged struct. See
// https://godoc.org/github.com/google/go-querystring/query for details.
func (s *Client) QueryStruct(queryStruct interface{}) *Client {
	if queryStruct != nil {
		s.queryStructs = append(s.queryStructs, queryStruct)
	}
	return s
}

// Body sets the Client's body. The body value will be set as the Body on new
// requests (see Request()).
// If the provided body is also an io.Closer, the request Body will be closed
// by http.Client methods.
func (s *Client) Body(body io.Reader) *Client {
	if body == nil {
		return s
	}
	return s.BodyProvider(bodyProvider{body: body})
}

// BodyProvider sets the Client's body provider.
func (s *Client) BodyProvider(body BodyProvider) *Client {
	if body == nil {
		return s
	}
	s.bodyProvider = body

	ct := body.ContentType()
	if ct != "" {
		s.Set(contentType, ct)
	}

	return s
}

// BodyJSON sets the Client's bodyJSON. The value pointed to by the bodyJSON
// will be JSON encoded as the Body on new requests (see Request()).
// The bodyJSON argument should be a pointer to a JSON tagged struct. See
// https://golang.org/pkg/encoding/json/#MarshalIndent for details.
func (s *Client) BodyJSON(bodyJSON interface{}) *Client {
	if bodyJSON == nil {
		return s
	}
	return s.BodyProvider(jsonBodyProvider{payload: bodyJSON})
}

// BodyForm sets the Client's bodyForm. The value pointed to by the bodyForm
// will be url encoded as the Body on new requests (see Request()).
// The bodyForm argument should be a pointer to a url tagged struct. See
// https://godoc.org/github.com/google/go-querystring/query for details.
func (s *Client) BodyForm(bodyForm interface{}) *Client {
	if bodyForm == nil {
		return s
	}
	return s.BodyProvider(formBodyProvider{payload: bodyForm})
}

// Requests

// Request returns a new http.Request created with the Client properties.
// Returns any errors parsing the rawURL, encoding query structs, encoding
// the body, or creating the http.Request.
func (s *Client) Request(ctx context.Context) (*http.Request, error) {
	reqURL, err := url.Parse(s.rawURL)
	if err != nil {
		return nil, err
	}

	err = addQueryStructs(reqURL, s.queryStructs)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if s.bodyProvider != nil {
		body, err = s.bodyProvider.Body()
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, s.method, reqURL.String(), body)
	if err != nil {
		return nil, err
	}
	addHeaders(req, s.header)

	if s.signer != nil {
		err = s.signer.Sign(req)
		if err != nil {
			return nil, err
		}
	}

	return req, err
}

// addQueryStructs parses url tagged query structs using go-querystring to
// encode them to url.Values and format them onto the url.RawQuery. Any
// query parsing or encoding errors are returned.
func addQueryStructs(reqURL *url.URL, queryStructs []interface{}) error {
	urlValues, err := url.ParseQuery(reqURL.RawQuery)
	if err != nil {
		return err
	}
	// encodes query structs into a url.Values map and merges maps
	for _, queryStruct := range queryStructs {
		queryValues, err := goquery.Values(queryStruct)
		if err != nil {
			return err
		}
		for key, values := range queryValues {
			for _, value := range values {
				urlValues.Add(key, value)
			}
		}
	}
	// url.Values format to a sorted "url encoded" string, e.g. "key=val&foo=bar"
	reqURL.RawQuery = urlValues.Encode()
	return nil
}

// addHeaders adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func addHeaders(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

// Sending

// ResponseDecoder sets the Client's response decoder.
func (s *Client) ResponseDecoder(decoder ResponseDecoder) *Client {
	if decoder == nil {
		return s
	}
	s.responseDecoder = decoder
	return s
}

// ReceiveSuccess creates a new HTTP request and returns the response. Success
// responses (2XX) are JSON decoded into the value pointed to by successV.
// Any error creating the request, sending it, or decoding a 2XX response
// is returned.
func (s *Client) ReceiveSuccess(ctx context.Context, successV interface{}) (*http.Response, error) {
	return s.Receive(ctx, successV, nil)
}

// Receive creates a new HTTP request and returns the response. Success
// responses (2XX) are JSON decoded into the value pointed to by successV and
// other responses are JSON decoded into the value pointed to by failureV.
// If the status code of response is 204(no content) or the Content-Length is 0,
// decoding is skipped. Any error creating the request, sending it, or decoding
// the response is returned.
// Receive is shorthand for calling Request and Do.
func (s *Client) Receive(ctx context.Context, successV, failureV interface{}) (*http.Response, error) {
	req, err := s.Request(ctx)
	if err != nil {
		return nil, err
	}
	return s.Do(req, successV, failureV)
}

// Do sends an HTTP request and returns the response. Success responses (2XX)
// are JSON decoded into the value pointed to by successV and other responses
// are JSON decoded into the value pointed to by failureV.
// If the status code of response is 204(no content) or the Content-Length is 0,
// decoding is skipped. Any error sending the request or decoding the response
// is returned.
func (s *Client) Do(req *http.Request, successV, failureV interface{}) (*http.Response, error) {
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()

	// The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	// See: https://golang.org/pkg/net/http/#Response
	defer io.Copy(io.Discard, resp.Body)

	// Don't try to decode on 204s or Content-Length is 0
	if resp.StatusCode == http.StatusNoContent || resp.ContentLength == 0 {
		return resp, nil
	}

	// Decode from json
	if successV != nil || failureV != nil {
		err = decodeResponse(resp, s.responseDecoder, successV, failureV)
	}
	return resp, err
}

// decodeResponse decodes response Body into the value pointed to by successV
// if the response is a success (2XX) or into the value pointed to by failureV
// otherwise. If the successV or failureV argument to decode into is nil,
// decoding is skipped.
// Caller is responsible for closing the resp.Body.
func decodeResponse(resp *http.Response, decoder ResponseDecoder, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decoder.Decode(resp, successV)
		}
	} else {
		if failureV != nil {
			return decoder.Decode(resp, failureV)
		}
	}
	return nil
}

package botid

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tlsclient "github.com/bogdanfinn/tls-client"
)

type BotID struct {
	client    tlsclient.HttpClient
	scriptUrl string
}

func NewBotID(scriptUrl string) (*BotID, error) {
	jar := tlsclient.NewCookieJar()
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithClientProfile(Brave_144),
		tlsclient.WithCookieJar(jar),
	}

	client, err := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}
	client.SetFollowRedirect(true)
	return &BotID{client: client, scriptUrl: scriptUrl}, nil
}

func (bot *BotID) FetchScript() (*string, error) {
	req, err := http.NewRequest(http.MethodGet, bot.scriptUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"sec-ch-ua":                 {"\"Chromium\";v=\"142\", \"Brave\";v=\"142\", \"Not_A Brand\";v=\"99\""},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"\"Windows\""},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"},
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"},
		"sec-gpc":                   {"1"},
		"accept-language":           {"en-US,en;q=0.5"},
		"sec-fetch-site":            {"none"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-user":            {"?1"},
		"sec-fetch-dest":            {"document"},
		"accept-encoding":           {"gzip, deflate, br, zstd"},
		"priority":                  {"u=0, i"},
		http.HeaderOrderKey:         {"sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "upgrade-insecure-requests", "user-agent", "accept", "sec-gpc", "accept-language", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-user", "sec-fetch-dest", "accept-encoding", "priority"},
	}

	resp, err := bot.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(body)

	return &bodyString, nil
}

func (bot *BotID) Verify(token string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, "https://botid-testing-pi.vercel.app/api/contact/test", nil)
	if err != nil {
		return "", err
	}

	req.Header = http.Header{
		"content-length":     {"0"},
		"sec-ch-ua-platform": {"\"Windows\""},
		"x-is-human":         {token},
		"x-path":             {"/api/generate"},
		"user-agent":         {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"},
		"sec-ch-ua":          {"\"Chromium\";v=\"142\", \"Brave\";v=\"142\", \"Not_A Brand\";v=\"99\""},
		"x-method":           {"POST"},
		"sec-ch-ua-mobile":   {"?0"},
		"accept":             {"*/*"},
		"sec-gpc":            {"1"},
		"accept-language":    {"en-US,en;q=0.5"},
		"origin":             {"https://botid-testing-pi.vercel.app"},
		"sec-fetch-site":     {"same-origin"},
		"sec-fetch-mode":     {"cors"},
		"sec-fetch-dest":     {"empty"},
		"referer":            {"https://botid-testing-pi.vercel.app/"},
		"accept-encoding":    {"gzip, deflate, br, zstd"},
		"priority":           {"u=1, i"},
		http.HeaderOrderKey:  {"content-length", "sec-ch-ua-platform", "x-is-human", "x-path", "user-agent", "sec-ch-ua", "x-method", "sec-ch-ua-mobile", "accept", "sec-gpc", "accept-language", "origin", "sec-fetch-site", "sec-fetch-mode", "sec-fetch-dest", "referer", "accept-encoding", "priority"},
	}

	resp, err := bot.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (bot *BotID) GenerateToken() (string, error) {
	script, err := bot.FetchScript()
	if err != nil {
		return "", err
	}
	start := time.Now()
	ctx, err := ExtractFromScript(script)
	if err != nil {
		return "", err
	}

	payload, err := BuildPayload(ctx)
	if err != nil {
		return "", err
	}

	encoded, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	fmt.Println(time.Now().Sub(start))
	return string(encoded), nil
}

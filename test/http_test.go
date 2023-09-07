func TestGet(t *testing.T) {
    r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	indexHandler(w, r, nil)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println("response:", resp.StatusCode, resp.Header, string(body))

	if resp.StatusCode != 200 {
		fmt.Println("http status:", http.StatusText(resp.StatusCode))
		t.Fatal()
	}
}

func TestPost(t *testing.T) {
    form := url.Values{}
	form.Add("key", "value")
    r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	r.Form = form
    w := httptest.NewRecorder()
	indexHandler(w, r, nil)
    resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println("response:", resp.StatusCode, resp.Header, string(body))

	if resp.StatusCode != 200 {
		fmt.Println("http status:", http.StatusText(resp.StatusCode))
		t.Fatal()
	}

}

func TestPostMultipartForm(t *testing.T) {
    // credit: https://federico-arias.github.io/software/testing-upload-handler/
    pr, pw := io.Pipe()
	form := multipart.NewWriter(pw)
	go func() {
		form.WriteField("key", "value")
		form.Close()
	}()
    r := httptest.NewRequest("POST", "/", pr)
	r.Header.Add("Content-Type", form.FormDataContentType())
    w := httptest.NewRecorder()
	indexHandler(w, r, nil)
    resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	fmt.Println("response:", resp.StatusCode, resp.Header, string(body))

	if resp.StatusCode != 200 {
		fmt.Println("http status:", http.StatusText(resp.StatusCode))
		t.Fatal()
	}
}

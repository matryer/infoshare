package comments_test

// const url = "http://localhost:8080"
//
// type Comment struct {
// 	ID     string `json:"id"`
// 	Body   string `json:"body"`
// 	Author string `json:"author"`
// }

//
// func TestComments(t *testing.T) {
// 	h := comments.NewHandler()
// 	assert.NotNil(t, h)
//
// 	res, err := http.Post(url+"/comments", "application/json", strings.NewReader(`{
//     "url": "https://infoshare.pl/",
//     "author": "John",
//     "body": "I love this conference."
//   }`))
// 	assert.NoError(t, err, "http.Post failed")
// 	assert.Equal(t, res.StatusCode, http.StatusCreated)
//
// 	var newComment Comment
// 	err = json.NewDecoder(res.Body).Decode(&newComment)
// 	assert.NoError(t, err)
//
// 	assert.NotEmpty(t, newComment.ID)
// 	assert.Equal(t, newComment.Body, "I love this conference.")
// 	assert.Equal(t, newComment.Author, "John")
//
// 	res, err = http.Get(url + "/comments?url=https://infoshare.pl/")
// 	assert.NoError(t, err, "http.Get failed")
//
// 	var cs []Comment
// 	err = json.NewDecoder(res.Body).Decode(&cs)
// 	assert.NoError(t, err)
// 	assert.Equal(t, len(cs), 1)
// 	assert.Equal(t, cs[0].ID, newComment.ID)
// 	assert.Equal(t, cs[0].Body, "I love this conference.")
// 	assert.Equal(t, cs[0].Author, "John")
//
// }

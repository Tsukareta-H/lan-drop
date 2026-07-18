###知識點

- http.Dir()獲得的路徑會和http.Handle()中第一個参數的路徑拼接.所以用http.StripPrefix()删掉http.Handle()中的路徑
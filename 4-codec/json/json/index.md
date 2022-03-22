### 编码
- json.NewEncoder(<Writer>).encode(v)
- json.Marshal(&v)
### 解码
- json.NewDecoder(<Reader>).decode(&v)
- json.Unmarshal([]byte, &v)

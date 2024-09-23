source .env

echo $OPENAI_KEY

curl https://api.openai.com/v1/chat/completions \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $OPENAI_KEY"  \
-d '{
        "model": "gpt-4o",
        "messages": [
            {"role": "user", "content": "write a haiku about ai"}
        ]
    }'

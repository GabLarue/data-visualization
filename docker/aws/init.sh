awslocal s3api create-bucket --bucket bucket

{
    "Location": "/bucket"
}

awslocal s3api list-buckets

{
    "Buckets": [
        {
            "Name": "bucket",
            "CreationDate": "2022-09-19T10:48:38+00:00"
        }
    ],
    "Owner": {
        "DisplayName": "webfile",
        "ID": "bcaf1ffd86f41161ca5fb16fd081034f"
    }
}

awslocal s3api put-object --bucket bucket --key index.html --body index.html

{
    "ETag": "\"d41d8cd98f00b204e9800998ecf8427e\""
}
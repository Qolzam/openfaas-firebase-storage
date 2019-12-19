# OpenFaaS firebase storage using Golang Http middleware template


## Setup
> I assume you are running your OpenFaaS server on `http://127.0.0.1:31112`.
> If not please change the URL in code to your own OpenFaaS server URL.
1. Generate a [Firebase private key](https://firebase.google.com/docs/admin/setup#initialize-sdk) from firebase console and your project setting to initialize the SDK.
2. Copy the service account key file into `storage` folder with the name `serviceAccountKey.json`.
3. Change the bucket name(`bucket_name`) in `stack.yml` file environment variable to your firebase bucket name.
4. Build and deploy the function:
```bash
faas up
```

## Example
You can test your function with [upload.html](https://github.com/Qolzam/openfaas-firebase-storage/blob/master/upload.html) file. In response you will have the url to access to your file.




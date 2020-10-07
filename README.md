# Script to test Google Search Console API

## 🎯 Features

- List sitemaps
- List sites that user have access
- Search analytics for a specific site URL

## ⚙️ Create a private key to access the GCP

For this method, you'll need to [create a service account](https://cloud.google.com/docs/authentication/getting-started), and download a key.

1. In the GCP Console, go to the [Create service account key](https://console.cloud.google.com/apis/credentials/serviceaccountkey?_ga=2.44822625.-475179053.1491320180) page;
2. From the Service account drop-down list, select New service account;
3. In the Service account name field, enter a name;
4. From the Role drop-down list, select Project > Owner or if you wanna more restrict roles you can choose: `cloudfunctions.admin` (Cloud Functions Admin) and `iam.serviceAccountUser` (Service Account User);
   1. The service account needs to Cloud Functions Admin to have permission to update the IAM policies;
5. Click Create. A JSON file that contains your key downloads to your computer;
6. Save the JSON file to a safe place on your local machine (I create a hidden folder and save it in `~/google-cloud-keys/key.json` );
7. Update the constant `projectID` with the name of your project at `main.go`;

## Run

```bash
GOOGLE_APPLICATION_CREDENTIALS=~/google-cloud-keys/key.json go run *.go 
```
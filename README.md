# Automation Project

Follow these steps to set up the project:

1. Setup Command-Line Interfaces  
   - Install **gcloud CLI** for managing Google Cloud resources  
     [Installation Guide](https://cloud.google.com/sdk/docs/install)  
   - Install **Pulumi CLI** for deploying and managing infrastructure  
     [Installation Guide](https://www.pulumi.com/docs/get-started/install/)

2. Configure Authentication  
   - Initialize gcloud CLI by running:
     ```bash
     gcloud init
     ```
     and connect with your Gmail account, then select your newly created GCP project.
   - Then, authenticate with:
     ```bash
     gcloud auth application-default login
     ```
   - If prompted for quota configuration, run:
     ```bash
     gcloud auth application-default set-quota-project <project-id>
     ```

3. Store API Tokens Securely  
   You can securely store these tokens using Pulumi by adding them as "secret" configuration variables. For example, after obtaining your tokens, store them locally by running:
   ```bash
   pulumi config set --secret gmail:accessToken <your_access_token>
   pulumi config set --secret gmail:refreshToken <your_refresh_token>
   ```
   These values are then stored in Pulumi's state in an encrypted form, and we retrieve them by using os.Getenv via Pulumi or through the Pulumi configuration package.
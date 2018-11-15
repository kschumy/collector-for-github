## About
An AWS Lambda to query GitHub and post results


## Setup
1. Create an S3 bucket.
    1. The bucket must be private, but your AWS account must have read/write permission. 
    1. The bucket name will later be the env variable `BUCKET_NAME` and the region will be `AWS_REGION` (but this should be completed for you by the Lambda and should not be manually added to the evn vars).
1. Upload the `config.yaml` file to the bucket. You can change the file name if you choose.
    1. Your AWS account must have read/write permission to the file.
    1. The file name will later be the env variable `CONFIG_FILE_NAME`.
1. Using a GitHub account, make a token. This token will be used for querying GitHub. 
    1. All permission (expect gist, delete_repo, and those for hooks) are required. Although not required, it is **highly** recommended that the GitHub account associated with this token should not be used for any other API calls.
    1. The token will later be the env variable `GITHUB_GRAPHQL_TOKEN`.
1. Using a **different GitHub account**, make a second token for posting to the team discussion board on GitHub
    1. All permission (expect gist, delete_repo, and those for hooks) are required. The GitHub account associated with this token **MUST** be able to post issues to the team's discussion board.
    1. The token will later be the env variable `GITHUB_TOKEN`.
1. Run the following code in your terminal. The output will be the env variable `BOARD_ID`.
``` bash
$ cmd=$(curl -s --request POST --url 'https://api.github.com/graphql?=' --header 'accept: application/vnd.github.starfire-preview+json' --header 'authorization: Bearer <GITHUB_TOKEN value>' --header 'content-type: application/json' --data '{"query":"query{\n  organization(login:\"<ORGANIZATION NAME>\") {\n    team(slug:\"<TEAM NAME>\") {\n      id\n\t\t}\n  }\n}"}' | jq -r '.data.organization.team.id' | base64 --d) 
$ echo ${cmd#*Team}
```
1. Create role with the following permissions:
    1. AWSLambdaFullAccess
    1. AmazonS3FullAccess
    1. CloudWatchEventsFullAccess
1. Create a Lambda.
    1. Set the language to Go.
1. Connect the bucket to the Lambda
1. Create CloudWatch event that triggers every 90 minutes and connect it to the lambda
1. Upload the zipped program to Lambda
1. Add the above-listed env vars to the lambda

The lambda may report failing the first time(s) it is run. Check the logs for the reason. If there are more than 120 issues, it will do them in batched but will "fail" between runs until it has processed all the issues. Check the logs for more information if you get errors.

## Known Issues
- Throttling on posting issues/managing the Lambda time limit is very poorly implemented. 
  - Getting all issues but only posting 20 doesn't make sense. The package the program uses will eventually allow for the user to specify how many should be queried for, and this program will use this feature in the future.
  - Force fail and banking on the Lambda to restart is not a great way to deal with processing results.
- Does not go back and re-search for updated issues that later meet the search criteria. There are no current plans to address this.

## Possible Features in the Future
- Connect to Wiki
- If/when GitHub allows GraphQL to make discussion board posts, the program is planning on switching to that.
- Manage terms and labels from an external yaml file

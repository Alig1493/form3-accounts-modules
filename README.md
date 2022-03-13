# from3-accounts-modules
A go module created using form3 acccounts apis.

## Candidate Profile:
- Name: Mohammed Ali Zubair
- Experience with go: Relevant working experience is within the timeline from the start to finish of this project. SO 2 weeks.
- Backend Engineering experience of 5 years.

## Instructions
- To run the project a simple `docker-compose up` and the `go` service will simply run the tests available in the module and then exit with a 0 status.
- To run just the go service `docker-compose run go`

## Work Considerations:
- I used both mocking and actual requests to test the module with the fake API.
- I took the approach which I thought was best and it can be made much better and DRY in my opinion after getting considerable experience in this new language
- I wrote the main file initially to test and nothing more.
- Probably would have liked to write a lot more test scenarios but I would also like to submit this in a respectable time as well. This is also considering the fact that I just moved to a new place and I am having to do a lot of other things for the time being besides software development.

## Submission Guidance

### Checklist

The finished solution **should:**
- [x] Be written in Go.
- [x] Use the `docker-compose.yaml` of this repository.
- [x] Be a client library suitable for use in another software project.
- [x] Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- [x] Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API. (There are lackings in the kinds of scenarios but I hope to do better with time.)
- [x] Be simple and concise. (Could have been simpler I feel and I hope to reach that level in due time as well.)
- [x] Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.
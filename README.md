# ghost-wp
just a wordpack since i couldn't find any around

wanted a ghost game but not sure what words to go for? 

## Contributing

please add more suggestions to the word pack ^^

## Usage

You can directly use the `easy` file, or make an API call to 

`ghost-wp.netlify.app`

## Deploying to Heroku

Set the buildpack using Heroku-CLI

`heroku login`

`heroku create -b https://github.com/kr/heroku-buildpack-go.git`

`git push heroku master`

To set the buildpack after create, do

`heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go`

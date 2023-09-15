# Running the Docusaurus-Generator in Local

## prerequisites
* [Node.js](https://nodejs.org/en/download/)
* [Yeoman](https://yeoman.io/learning/index.html)
### Instructions

Install project dependencies and symlink a global module to the local file.
```
npm link
```

Run Yeoman Generator to scaffold a sample Docusaurus application
```
yo docusaurus
```
If you wish to give custom name for your Project then run the command as
```
yo docusaurus --file inputfile.json
```
For giving input add your .json file to the generator directory. Example:-
```
{
    "applicationName": "Test_gen111",
    "serverPort": 5000
}
```

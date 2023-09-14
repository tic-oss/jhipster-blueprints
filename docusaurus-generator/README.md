# Running the Docusaurus-Generator in Local

## prerequisites
* [Node.js](https://nodejs.org/en/download/)
* [Yeoman](https://yeoman.io/learning/index.html)
* [Maven](https://maven.apache.org/download.cgi)
* JDK(17)
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
yo docusaurus --projectName your_project_name
```

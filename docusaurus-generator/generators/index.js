const fs = require("fs");
const path = require("path");
const Generator = require("yeoman-generator");

module.exports = class extends Generator {
  constructor(args, opts) {
    super(args, opts);

    this.argument("projectName", {
      type: String,
      required: false,
      description: "Name of the project"
    });

    this.option("generateDocusaurus", {
      description: "Generate Docusaurus documentation",
      type: Boolean,
      default: true
    });
  }

  writing() {
    const projectName = this.options.projectName || "MyDocumentation"; // Default project name
    const copyOpts = {
      globOptions: {
        ignore: []
      }
    };

    const options = {
      projectName: projectName
    };

    if (this.options.generateDocusaurus) {
      this._generateDocusaurus(options, copyOpts);
    }

    this._generateOtherFiles(options, copyOpts);
  }

  _generateDocusaurus(options, copyOpts) {
    this.fs.copyTpl(
      this.templatePath(`docusaurus`), 
      this.destinationPath(`docusaurus-${options.projectName}`),
      options,
      copyOpts
    );
  
    // Read the template docusaurus.config.js file
    const templateConfigPath = this.templatePath("docusaurus/docusaurus.config.js");
    const templateConfigContent = this.fs.read(templateConfigPath);
  
    // Perform string replacement for the project name
    const updatedConfigContent = templateConfigContent.replace(
      /<%= projectName %>/g,
      options.projectName
    );
  
    // Write the updated config to the generated directory
    const generatedConfigPath = this.destinationPath(
      `docusaurus-${options.projectName}/docusaurus.config.js`
    );
    this.fs.write(generatedConfigPath, updatedConfigContent);
  }
  
  
  _generateOtherFiles(options, copyOpts) {
    // Generate other files logic
    const filesToGenerate = [
        "docs/intro.md",
        "docs/Documentation/concept.md",
        "docs/Documentation/maintopic.md",
        "docs/Documentation/subfolder/subfile.md",
        //"blog/2019-05-28-five-blog-post.md",
        "blog/2021-08-01-mdx-blog-post.mdx",
        "blog/2023-08-29-three-blog-post.md",
        "blog/2023-08-29-four-blog-post.md",
        "blog/2023-08-29-six-blog-post.md",
        "blog/2023-08-29-seven-blog-post.md",
        "blog/2023-08-29-eight-blog-post.md",
        "blog/authors.yml",
        "src/components/HomepageFeatures/index.js",
        "src/components/HomepageFeatures/styles.module.css",
        "src/css/custom.css",
        "src/pages/index.js",
        "src/pages/index.module.css",
        "src/theme/BlogListPage/Author/index.js",
        "src/theme/BlogListPage/Author/styles.module.css",
        "src/theme/BlogListPage/ListItem/index.js",
        "src/theme/BlogListPage/ListItem/styles.module.css",
        "src/theme/BlogListPage/index.js",
        "src/theme/BlogListPage/styles.module.css",
        "static/img/image.jpeg",
        "static/img/logo.png",
        "docusaurus.config.js",
        "sidebars.js",
        "package.json",
        "README.md"
      ];

    filesToGenerate.forEach(file => {
      this.fs.copyTpl(
        this.templatePath(`docusaurus/${file}`), // Add a forward slash here
        this.destinationPath(`docusaurus-${options.projectName}/${file}`),
        options,
        copyOpts
      );
    });
  }

  // Other helper methods, prompts, and install logic
};

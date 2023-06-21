import chalk from "chalk";
import yosay from 'yosay';
import ServerGenerator from "generator-jhipster/generators/server";
import { askForServerSideOpts } from './prompts.mjs';

export default class extends ServerGenerator {
  constructor(args, opts, features) {
    super(args, opts, features);

    if (this.options.help) return;

    if (!this.options.jhipsterContext) {
      throw new Error(
        `This is a JHipster blueprint and should be used only like ${chalk.yellow(
          "jhipster --blueprints test"
        )}`
      );
    }
  }

  get [ServerGenerator.INITIALIZING]() {
    return {
      // ...super.initializing,
      // async initializingTemplateTask() {},
    };
  }

  get [ServerGenerator.PROMPTING]() {
    return {
      // ...super.prompting,
      // async promptingTemplateTask() {},
      prompting() {
        // Have Yeoman greet the user.
        this.log(
          yosay(
            `${chalk.red('gomicro-blueprint')}`
          )
        );
      },
      askForServerSideOpts
    };
  }

  get [ServerGenerator.CONFIGURING]() {
    return {
      // ...super.configuring,
      // async configuringTemplateTask() {},
    };
  }

  get [ServerGenerator.COMPOSING]() {
    return {
      // ...super.composing,
      // async composingTemplateTask() {},
    };
  }

  get [ServerGenerator.LOADING]() {
    return {
      // ...super.loading,
      // async loadingTemplateTask() {},
    };
  }

  get [ServerGenerator.PREPARING]() {
    return {
      // ...super.preparing,
      // async preparingTemplateTask() {},
    };
  }

  get [ServerGenerator.CONFIGURING_EACH_ENTITY]() {
    return {
      // ...super.configuringEachEntity,
      // async configuringEachEntityTemplateTask() {},
    };
  }

  get [ServerGenerator.LOADING_ENTITIES]() {
    return {
      // ...super.loadingEntities,
      // async loadingEntitiesTemplateTask() {},
    };
  }

  get [ServerGenerator.PREPARING_EACH_ENTITY]() {
    return {
      // ...super.preparingEachEntity,
      // async preparingEachEntityTemplateTask() {},
    };
  }

  get [ServerGenerator.PREPARING_EACH_ENTITY_FIELD]() {
    return {
      // ...super.preparingEachEntityField,
      // async preparingEachEntityFieldTemplateTask() {},
    };
  }

  get [ServerGenerator.PREPARING_EACH_ENTITY_RELATIONSHIP]() {
    return {
      // ...super.preparingEachEntityRelationship,
      // async preparingEachEntityRelationshipTemplateTask() {},
    };
  }

  get [ServerGenerator.POST_PREPARING_EACH_ENTITY]() {
    return {
      // ...super.postPreparingEachEntity,
      // async postPreparingEachEntityTemplateTask() {},
    };
  }

  get [ServerGenerator.DEFAULT]() {
    return {
      // ...super.default,
      // async defaultTemplateTask() {},
    };
  }

  get [ServerGenerator.WRITING]() {
    return {
      writing() {
        const templateVariables = {
          serverPort: this.serverPort,
          packageName: this.packageName,
          baseName: this.baseName,
          auth: this.auth,
          eureka: this.eureka,
          rabbitmq: this.rabbitmq,
          postgresql: this.postgress,
          mongodb: this.mongodb
        };
      
        const templatePaths = [
          { src: "gomicro/docker", dest: "docker" },
          { src: "gomicro/gomicro/proto", dest: "gomicro/proto" },
          { src: "gomicro/gomicro/go.mod", dest: "gomicro/go.mod" },
          { src: "gomicro/gomicro/main.go", dest: "gomicro/main.go" },
          { src: "gomicro/gomicro/Dockerfile", dest: "gomicro/Dockerfile" },
          { src: "gomicro/gomicro/Makefile", dest: "gomicro/Makefile" },
          { src: "gomicro/gomicro/README.md", dest: "gomicro/README.md" },
          { src: "gomicro/gomicro/.env", dest: "gomicro/.env" }
        ];
      
        const conditionalTemplates = [
          { condition: this.auth, src: "gomicro/gomicro/auth", dest: "gomicro/auth" },
          { condition: this.postgress || this.mongodb, src: "gomicro/gomicro/handler", dest: "gomicro/handler" },
          { condition: this.postgress || this.mongodb, src: "gomicro/gomicro/db", dest: "gomicro/db" },
          { condition: this.eureka, src: "gomicro/gomicro/eurekaregistry", dest: "gomicro/eurekaregistry" },
          { condition: this.rabbitmq, src: "gomicro/gomicro/rabbitmq", dest: "gomicro/rabbitmq" },
        ];
      
        templatePaths.forEach(({ src, dest }) => {
          this.fs.copyTpl(
            this.templatePath(src),
            this.destinationPath(dest),
            templateVariables
          );
        });
      
        conditionalTemplates.forEach(({ condition, src, dest }) => {
          if (condition) {
            this.fs.copyTpl(
              this.templatePath(src),
              this.destinationPath(dest),
              templateVariables
            );
          }
        });
      }
     }
}

  get [ServerGenerator.WRITING_ENTITIES]() {
    return {
      // ...super.writingEntities,
      // async writingEntitiesTemplateTask() {},
    };
  }

  get [ServerGenerator.POST_WRITING]() {
    return {
      // ...super.postWriting,
      // async postWritingTemplateTask() {},
    };
  }

  get [ServerGenerator.POST_WRITING_ENTITIES]() {
    return {
      // ...super.postWritingEntities,
      // async postWritingEntitiesTemplateTask() {},
    };
  }

  get [ServerGenerator.INSTALL]() {
    return {
      // ...super.install,
      // async installTemplateTask() {},
    };
  }

  get [ServerGenerator.POST_INSTALL]() {
    return {
      // ...super.postInstall,
      // async postInstallTemplateTask() {},
    };
  }

  get [ServerGenerator.END]() {
    return {
      // ...super.end,
      // async endTemplateTask() {},
    };
  }
}

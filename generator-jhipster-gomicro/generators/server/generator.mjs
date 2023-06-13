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
            `${chalk.red('golang-blueprint')}`
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
      // ...super.writing,
      // async writingTemplateTask() {
      //   await this.writeFiles({
      //     sections: {
      //       files: [{ templates: ["template-file-server"] }],
      //     },
      //     context: this,
      //   });
      // },
      writing() {
        console.log(this.serverPort,this.packageName,this.auth,this.eureka,this.rabbitmq);
        this.fs.copyTpl(
        this.templatePath("go/docker"),
        this.destinationPath("docker"), {
        serverPort: this.serverPort,
        packageName: this.packageName,
        baseName: this.baseName,
        auth:this.auth,
        eureka:this.eureka,
        rabbitmq:this.rabbitmq,
        postgresql:this.postgress
        }
        );
        if(this.auth){
        this.fs.copyTpl(
          this.templatePath("go/go/auth"),
          this.destinationPath("go/auth"), {
          serverPort: this.serverPort,
          packageName: this.packageName,
          baseName: this.baseName,
          auth:this.auth,
          eureka:this.eureka,
          rabbitmq:this.rabbitmq,
          postgresql:this.postgress
        }
        );
        }
        if(this.postgress){
          this.fs.copyTpl(
            this.templatePath("go/go/handler"),
            this.destinationPath("go/handler"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
          }
          );
          this.fs.copyTpl(
            this.templatePath("go/go/pkg"),
            this.destinationPath("go/pkg"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
          }
          );
          this.fs.copyTpl(
            this.templatePath("go/go/proto"),
            this.destinationPath("go/proto"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
          }
          );
        }
        this.fs.copyTpl(
          this.templatePath("go/go/go.mod"),
          this.destinationPath("go/go.mod"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
        }
        );
        this.fs.copyTpl(
          this.templatePath("go/go/main.go"),
          this.destinationPath("go/main.go"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
        }
        );
        this.fs.copyTpl(
          this.templatePath("go/go/Makefile"),
          this.destinationPath("go/Makefile"), {
          serverPort: this.serverPort
        }
        );
        this.fs.copyTpl(
          this.templatePath("go/go/README.md"),
          this.destinationPath("go/README.md"), {
          serverPort: this.serverPort
        }
        );
        this.fs.copyTpl(
          this.templatePath("go/go/.env"),
          this.destinationPath("go/.env"), {
            serverPort: this.serverPort,
            packageName: this.packageName,
            baseName: this.baseName,
            auth:this.auth,
            eureka:this.eureka,
            rabbitmq:this.rabbitmq,
            postgresql:this.postgress
        }
        );
      }
    };
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

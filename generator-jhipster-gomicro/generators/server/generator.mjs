import chalk from 'chalk';
import yosay from 'yosay';
import ServerGenerator from 'generator-jhipster/generators/server';
import { askForServerSideOpts } from './prompts.mjs';
import { loadCommunicationConfigs, findConfigByBaseName, deleteUnwantedFiles } from './util.mjs';

export default class extends ServerGenerator {
  constructor(args, opts, features) {
    super(args, opts, features);

    if (this.options.help) return;

    if (!this.options.jhipsterContext) {
      throw new Error(`This is a JHipster blueprint and should be used only like ${chalk.yellow('jhipster --blueprints test')}`);
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
        this.log(yosay(`${chalk.red('gomicro-blueprint')}`));
      },
      askForServerSideOpts,
      loadCommunicationConfigs,
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
      endTemplateTask() {
        deleteUnwantedFiles.call(this);
      },
      writing() {
        const matchingScenarios = findConfigByBaseName(this.baseName);

        if (matchingScenarios.length > 0) {
          var restServer = [],
            restClient,
            rabbitmqServer = [],
            rabbitmqClient = [];

          for (var options of matchingScenarios) {
            if (options.framework === 'rest-api') {
              if (options.server) {
                restServer.push(options.server);
              }
              if (options.client) {
                restClient = options.client;
              }
            } else if (options.framework === 'rabbitmq') {
              if (options.server) {
                rabbitmqServer.push(options.server);
              }
              if (options.client) {
                rabbitmqClient.push(options.client);
              }
            }
          }
        }
        const templateVariables = {
          serverPort: this.serverPort,
          packageName: this.packageName,
          baseName: this.baseName,
          auth: this.auth,
          eureka: this.eureka,
          rabbitmq: this.rabbitmq,
          postgresql: this.postgress,
          mongodb: this.mongodb,
          restServer: restServer,
          restClient: restClient,
          rabbitmqServer: rabbitmqServer,
          rabbitmqClient: rabbitmqClient,
        };

        const templatePaths = [
          { src: 'gomicro/docker', dest: 'docker' },
          { src: 'gomicro/gomicro/proto', dest: 'gomicro/proto' },
          { src: 'gomicro/gomicro/go.mod', dest: 'gomicro/go.mod' },
          { src: 'gomicro/gomicro/main.go', dest: 'gomicro/main.go' },
          { src: 'gomicro/gomicro/Dockerfile', dest: 'gomicro/Dockerfile' },
          { src: 'gomicro/gomicro/Makefile', dest: 'gomicro/Makefile' },
          { src: 'gomicro/gomicro/README.md', dest: 'gomicro/README.md' },
          { src: 'gomicro/gomicro/config', dest: 'gomicro/config' },
          { src: 'gomicro/gomicro/resources', dest: 'gomicro/resources' },
          { src: 'gomicro/gomicro/controllers', dest: 'gomicro/controllers' },
        ];
        const conditionalTemplates = [
          { condition: this.auth, src: 'gomicro/gomicro/auth', dest: 'gomicro/auth' },
          { condition: this.postgress, src: 'gomicro/gomicro/handler/db.go', dest: 'gomicro/handler/db.go' },
          { condition: this.mongodb, src: 'gomicro/gomicro/handler/mongodb.go', dest: 'gomicro/handler/mongodb.go' },
          { condition: this.postgress, src: 'gomicro/gomicro/db/config.go', dest: 'gomicro/db/config.go' },
          { condition: this.mongodb, src: 'gomicro/gomicro/db/mongoconfig.go', dest: 'gomicro/db/mongoconfig.go' },
          { condition: this.eureka, src: 'gomicro/gomicro/eurekaregistry/helper', dest: 'gomicro/eurekaregistry/helper' },
          {
            condition: this.eureka,
            src: 'gomicro/gomicro/eurekaregistry/DiscoveryManager.go',
            dest: 'gomicro/eurekaregistry/DiscoveryManager.go',
          },
          {
            condition: this.eureka,
            src: 'gomicro/gomicro/eurekaregistry/RegistrationManager.go',
            dest: 'gomicro/eurekaregistry/RegistrationManager.go',
          },
          {
            condition: this.eureka,
            src: 'gomicro/gomicro/eurekaregistry/EurekaRegistrationManager.go',
            dest: 'gomicro/eurekaregistry/EurekaRegistrationManager.go',
          },
          { condition: this.rabbitmq, src: 'gomicro/gomicro/rabbitmq', dest: 'gomicro/rabbitmq' },
          {
            condition: restServer?.length,
            src: 'gomicro/gomicro/eurekaregistry/ServiceDiscovery.go',
            dest: 'gomicro/eurekaregistry/ServiceDiscovery.go',
          },
          { condition: this.postgress, src: 'gomicro/gomicro/migrate', dest: 'gomicro/migrate' },
        ];
        templatePaths.forEach(({ src, dest }) => {
          this.fs.copyTpl(this.templatePath(src), this.destinationPath(dest), templateVariables);
        });
        conditionalTemplates.forEach(({ condition, src, dest }) => {
          if (condition) {
            this.fs.copyTpl(this.templatePath(src), this.destinationPath(dest), templateVariables);
          }
        });
        if (rabbitmqServer?.length) {
          for (var i = 0; i < rabbitmqServer.length; i++) {
            var server = rabbitmqServer[i].charAt(0).toUpperCase() + rabbitmqServer[i].slice(1);
            var client = this.baseName.charAt(0).toUpperCase() + this.baseName.slice(1);
            this.fs.copyTpl(
              this.templatePath('gomicro/gomicro/rabbitmq/consumer.go'),
              this.destinationPath('gomicro/rabbitmq/' + 'RabbitMQConsumer' + server + 'To' + client + '.go'),
              {
                packageName: this.packageName,
                rabbitmqServer: server,
                rabbitmqClient: client,
                baseName: this.baseName,
              }
            );
          }
        }
        if (rabbitmqClient?.length) {
          for (var i = 0; i < rabbitmqClient.length; i++) {
            var server = this.baseName.charAt(0).toUpperCase() + this.baseName.slice(1);
            var client = rabbitmqClient[i].charAt(0).toUpperCase() + rabbitmqClient[i].slice(1);
            this.fs.copyTpl(
              this.templatePath('gomicro/gomicro/rabbitmq/producer.go'),
              this.destinationPath('gomicro/rabbitmq/' + 'RabbitMQProducer' + server + 'To' + client + '.go'),
              {
                packageName: this.packageName,
                rabbitmqClient: client,
                rabbitmqServer: server,
                baseName: this.baseName,
              }
            );
          }
        }
      },
    };
  }
}

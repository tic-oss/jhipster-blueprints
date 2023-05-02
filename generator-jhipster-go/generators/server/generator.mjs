import chalk from 'chalk';
import yosay from 'yosay';
import ServerGenerator from 'generator-jhipster/generators/server';
// import {
//   PRIORITY_PREFIX,
//   INITIALIZING_PRIORITY,
//   PROMPTING_PRIORITY,
//   CONFIGURING_PRIORITY,
//   COMPOSING_PRIORITY,
//   LOADING_PRIORITY,
//   PREPARING_PRIORITY,
//   DEFAULT_PRIORITY,
//   WRITING_PRIORITY,
//   POST_WRITING_PRIORITY,
//   INSTALL_PRIORITY,
//   POST_INSTALL_PRIORITY,
//   END_PRIORITY,
// } from 'generator-jhipster/esm/priorities';

import { askForServerSideOpts } from './prompts.mjs';
export default class extends ServerGenerator {
  constructor(args, opts, features) {
    // super(args, opts, { taskPrefix: PRIORITY_PREFIX, ...features });
    super(args, opts, features);
    if (this.options.help) return;

    if (!this.options.jhipsterContext) {
      throw new Error(`This is a JHipster blueprint and should be used only like ${chalk.yellow('jhipster --blueprints go')}`);
    }
  }

  get [ServerGenerator.INITIALIZING]() {
    return {
      ...super._initializing(),
      async initializingTemplateTask() { },
    };
  }

  get [ServerGenerator.PROMPTING]() {
    return {

      // ...super._prompting(),
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
      // ...super._configuring(),
      // async configuringTemplateTask() {},
    };
  }

  get [ServerGenerator.COMPOSING]() {
    return {
      // ...super._composing(),
      // async composingTemplateTask() {},
    };
  }

  get [ServerGenerator.LOADING]() {
    return {
      // ...super._loading(),
      // async loadingTemplateTask() {},
    };
  }

  get [ServerGenerator.PREPARING]() {
    return {
      //   ...super._preparing(),
      //   async preparingTemplateTask() {},
    };
  }

  get [ServerGenerator.DEFAULT]() {
    return {
      // ...super._default(),
      // async defaultTemplateTask() {},
    };
  }

  get [ServerGenerator.WRITING]() {
    return {
      // ...super._writing(),
      // async writingTemplateTask() {
      //   await this.writeFiles({
      //     sections: {
      //       files: [{ templates: ['template-file-server'] }],
      //     },
      //     context: this,
      // });
      //},
      writing() {
        this.fs.copyTpl(
          this.templatePath("go"),
          this.destinationPath(), {
          serverPort: this.serverPort,
          packageName: this.packageName,
          baseName: this.baseName
        }
        );
        this.fs.copyTpl(
          this.templatePath("go/go/.env"),
          this.destinationPath("go/.env"), {
          serverPort: this.serverPort
        }
        );
      }
    };
  }

  get [ServerGenerator.POST_WRITING]() {
    return {
      ...super._postWriting(),
      async postWritingTemplateTask() { },
    };
  }

  get [ServerGenerator.INSTALL]() {
    return {
      ...super._install(),
      async installTemplateTask() { },
    };
  }

  get [ServerGenerator.POST_INSTALL]() {
    return {
      ...super._postInstall(),
      async postInstallTemplateTask() { },
    };
  }

  get [ServerGenerator.END]() {
    return {
      ...super._end(),
      async endTemplateTask() { },
    };
  }
}

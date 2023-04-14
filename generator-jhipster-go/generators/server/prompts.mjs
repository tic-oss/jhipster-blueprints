import chalk from 'chalk';
import _ from 'lodash';

// import {applicationTypes} from './types/application-types.js';
// import {databaseTypes} from './types/database-types.js';
// import {authenticationTypes} from './types/authentication-types.js';
// import {serviceDiscoveryTypes} from './types/service-discovery-types.js';

const APPLICATION_TYPE_MONOLITH = 'monolith';
const APPLICATION_TYPE_MICROSERVICE = 'microservice';
const APPLICATION_TYPE_GATEWAY = 'gateway';

const applicationTypes = {
  MONOLITH: APPLICATION_TYPE_MONOLITH,
  MICROSERVICE: APPLICATION_TYPE_MICROSERVICE,
  GATEWAY: APPLICATION_TYPE_GATEWAY,
};

const serviceDiscoveryTypes = {
  EUREKA: 'eureka',
  CONSUL: 'consul',
  NO_SERVICE_DISCOVERY: 'No service discovery',
};

const authenticationTypes = {
  JWT: 'jwt',
  OAUTH2: 'oauth2',
  SESSION: 'session',
};

const SQL = 'sql';
const MYSQL = 'mysql';
const MARIADB = 'mariadb';
const POSTGRESQL = 'postgresql';
const MSSQL = 'mssql';
const ORACLE = 'oracle';
const H2_DISK = 'h2Disk';
const H2_MEMORY = 'h2Memory';
const MONGODB= 'mongodb';
const CASSANDRA= 'cassandra';
const COUCHBASE= 'couchbase';
const NEO4J= 'neo4j';
const NO_DATABASE= 'No database';
const databaseTypes = {
  SQL,
  MYSQL,
  MARIADB,
  POSTGRESQL,
  MSSQL,
  ORACLE,
  MONGODB: 'mongodb',
  CASSANDRA: 'cassandra',
  COUCHBASE: 'couchbase',
  NEO4J: 'neo4j',
  H2_DISK,
  H2_MEMORY,
  NO_DATABASE: 'No database',
};
databaseTypes.isSql = type => [SQL, MYSQL, POSTGRESQL, ORACLE, MARIADB, MSSQL, H2_DISK, H2_MEMORY].includes(type);


// import { R2DBC_DB_OPTIONS, SQL_DB_OPTIONS } from './support/database.mjs';
const R2DBC_DB_OPTIONS = [
  {
    value: databaseTypes.POSTGRESQL,
    name: 'PostgreSQL',
  },
  {
    value: databaseTypes.MYSQL,
    name: 'MySQL',
  },
  {
    value: databaseTypes.MARIADB,
    name: 'MariaDB',
  },
  {
    value: databaseTypes.MSSQL,
    name: 'Microsoft SQL Server',
  },
];

const SQL_DB_OPTIONS = [
  {
    value: databaseTypes.POSTGRESQL,
    name: 'PostgreSQL',
  },
  {
    value: databaseTypes.MYSQL,
    name: 'MySQL',
  },
  {
    value: databaseTypes.MARIADB,
    name: 'MariaDB',
  },
  {
    value: databaseTypes.ORACLE,
    name: 'Oracle',
  },
  {
    value: databaseTypes.MSSQL,
    name: 'Microsoft SQL Server',
  },
];


const { GATEWAY, MICROSERVICE, MONOLITH } = applicationTypes;
const { JWT, OAUTH2, SESSION } = authenticationTypes;
// const { CASSANDRA, H2_DISK, H2_MEMORY, MONGODB, NEO4J, SQL, COUCHBASE } = databaseTypes;
const { CONSUL, EUREKA,NO_SERVICE_DISCOVERY } = serviceDiscoveryTypes;

const OptionNames = {
  AUTHENTICATION_TYPE: 'authenticationType',
  DATABASE_TYPE: 'databaseType',
  DEV_DATABASE_TYPE: 'devDatabaseType',
  PACKAGE_NAME: 'packageName',
  PROD_DATABASE_TYPE: 'prodDatabaseType',
  SERVER_PORT: 'serverPort',
  SERVICE_DISCOVERY_TYPE: 'serviceDiscoveryType'
};

const {
  AUTHENTICATION_TYPE,
  DATABASE_TYPE,
  DEV_DATABASE_TYPE,
  PROD_DATABASE_TYPE,
  PACKAGE_NAME,
  SERVER_PORT,
  SERVICE_DISCOVERY_TYPE,
} = OptionNames;


/**
 * Get Option From Array
 *
 * @param {Array} array - array
 * @param {any} option - options
 * @returns {boolean} true if option is in array and is set to 'true'
 */
const getOptionFromArray = (array, option) => {
  let optionValue = false;
  array.forEach(value => {
    if (_.includes(value, option)) {
      optionValue = value.split(':')[1];
    }
  });
  optionValue = optionValue === 'true' ? true : optionValue;
  return optionValue;
};

export async function askForServerSideOpts() {
<<<<<<< HEAD
  //console.log(this);
  if (this.options.fromJdl){ 
      this.serverPort=this.jhipsterConfigWithDefaults.serverPort;
      this.packageName=this.jhipsterConfigWithDefaults.packageName;
      this.baseName=this.jhipsterConfigWithDefaults.baseName;
    return ;
  }
  // if (!this.options.askAnswered) return;
  // if (!this.options.askAnswered){
  //   this.serverPort=this.jhipsterConfigWithDefaults.serverPort;
  //   this.packageName=this.jhipsterConfigWithDefaults.packageName;
  //   this.baseName=this.jhipsterConfigWithDefaults.baseName;
  //   return ;
  // }
/* answers structure   
   {
      serverPort: '8080',
      packageName: 'tic',
      serviceDiscoveryType: 'eureka',
      authenticationType: 'oauth2',
      databaseType: 'sql',
      prodDatabaseType: 'postgresql',
      devDatabaseType: 'postgresql'
    }
*/
    
  // const { applicationType, serverPort: defaultServerPort, reactive } = this.jhipsterConfigWithDefaults;

  //   const prompts = [
  //       {
  //         type: "string",
  //         name: "yourName",
  //         message: "What is the your name?",
  //         default: "tic"
  //       }
  //     ]
  //     return this.prompt(prompts).then(answers => {
  //       this.yourName = answers.yourName;
  //       // this.description = answers.description;
  //       // this.serverPort = answers.serverPort;
  //     });
  //   // await this.prompt(prompts, this.config);

=======
  // console.log(this.jhipsterConfigWithDefaults);
  // if (!this.options.askAnswered) return;
  if (!this.options.askAnswered){
    this.serverPort=this.jhipsterConfigWithDefaults.serverPort;
    this.packageName=this.jhipsterConfigWithDefaults.packageName;
    this.baseName=this.jhipsterConfigWithDefaults.baseName
    return ;
  }
>>>>>>> c61a02c5c30e40001d394063b0a0d61069d7d530

    const { applicationType, serverPort: defaultServerPort, reactive } = this.jhipsterConfigWithDefaults;
  const prompts = [
    {
      when: () => applicationType === GATEWAY || applicationType === MICROSERVICE,
      type: 'input',
      name: SERVER_PORT,
      validate: input => (/^([0-9]*)$/.test(input) ? true : 'This is not a valid port number.'),
      message:
        'As you are running in a microservice architecture, on which port would like your server to run? It should be unique to avoid port conflicts.',
      default: defaultServerPort,
    },
    {
      type: 'input',
      name: PACKAGE_NAME,
      validate: input =>
        /^([a-z_]{1}[a-z0-9_]*(\.[a-z_]{1}[a-z0-9_]*)*)$/.test(input)
          ? true
          : 'The package name you have provided is not a valid Java package name.',
      message: 'What is your default go package name?',
      default:  this.jhipsterConfigWithDefaults.packageName,
      store: true,
    },
    {
      when: () => applicationType === 'gateway' || applicationType === 'microservice',
      type: 'list',
      name: SERVICE_DISCOVERY_TYPE,
      message: 'Which service discovery server do you want to use?',
      choices: [
        {
          value: EUREKA,
          name: 'JHipster Registry (recommended)',
        },
        {
          value: CONSUL,
          name: 'Consul',
        },
        {
          value: NO_SERVICE_DISCOVERY,
          name: 'No service discovery',
        },
      ],
      default: this.jhipsterConfigWithDefaults.serviceDiscoveryType,
    },
    {
      when: answers =>
        (applicationType === MONOLITH && answers.serviceDiscoveryType !== EUREKA) || [GATEWAY, MICROSERVICE].includes(applicationType),
      type: 'list',
      name: AUTHENTICATION_TYPE,
      message: `Which ${chalk.yellow('*type*')} of authentication would you like to use?`,
      choices: answers => {
        const opts = [
          {
            value: JWT,
            name: 'JWT authentication (stateless, with a token)',
          },
        ];
        opts.push({
          value: OAUTH2,
          name: 'OAuth 2.0 / OIDC Authentication (stateful, works with Keycloak and Okta)',
        });
        if (applicationType === MONOLITH && answers.serviceDiscoveryType !== EUREKA) {
          opts.push({
            value: SESSION,
            name: 'HTTP Session Authentication (stateful, default Spring Security mechanism)',
          });
        }
        return opts;
      },
      default: this.jhipsterConfigWithDefaults.authenticationType,
    },
    {
      type: 'list',
      name: DATABASE_TYPE,
      message: `Which ${chalk.yellow('*type*')} of database would you like to use?`,
      choices: answers => {
        const opts = [];
        if (!answers.reactive) {
          opts.push({
            value: SQL,
            name: 'SQL (H2, PostgreSQL, MySQL, MariaDB, Oracle, MSSQL)',
          });
        } else {
          opts.push({
            value: SQL,
            name: 'SQL (H2, PostgreSQL, MySQL, MariaDB, MSSQL)',
          });
        }
        opts.push({
          value: MONGODB,
          name: 'MongoDB',
        });
        if (answers.authenticationType !== OAUTH2) {
          opts.push({
            value: CASSANDRA,
            name: 'Cassandra',
          });
        }
        opts.push({
          value: 'couchbase',
          name: '[BETA] Couchbase',
        });
        opts.push({
          value: NEO4J,
          name: '[BETA] Neo4j',
        });
        opts.push({
          value: NO_DATABASE,
          name: 'No database',
        });
        return opts;
      },
      default: this.jhipsterConfigWithDefaults.databaseType,
    },
    {
      when: response => response.databaseType === SQL,
      type: 'list',
      name: PROD_DATABASE_TYPE,
      message: `Which ${chalk.yellow('*production*')} database would you like to use?`,
      choices: answers => (answers.reactive ? R2DBC_DB_OPTIONS : SQL_DB_OPTIONS),
      default: this.jhipsterConfigWithDefaults.prodDatabaseType,
    },
    {
      when: response => response.databaseType === SQL,
      type: 'list',
      name: DEV_DATABASE_TYPE,
      message: `Which ${chalk.yellow('*development*')} database would you like to use?`,
      choices: response =>
        [
          {
            value: H2_DISK,
            name: 'H2 with disk-based persistence',
          },
          {
            value: H2_MEMORY,
            name: 'H2 with in-memory persistence',
          },
        ].concat(SQL_DB_OPTIONS.find(it => it.value === response.prodDatabaseType)),
      default: this.jhipsterConfigWithDefaults.devDatabaseType,
    },];

    return this.prompt(prompts).then(answers=>{
      this.serverPort=answers.serverPort;
      this.packageName=answers.packageName;
<<<<<<< HEAD
      this.baseName=this.jhipsterConfig.baseName;
=======
      this.baseName=this.jhipsterConfigWithDefaults.baseName
>>>>>>> c61a02c5c30e40001d394063b0a0d61069d7d530
  /* answers structure   
     {
        serverPort: '8080',
        packageName: 'tic',
        serviceDiscoveryType: 'eureka',
        authenticationType: 'oauth2',
        databaseType: 'sql',
        prodDatabaseType: 'postgresql',
        devDatabaseType: 'postgresql'
      }
  */
      
    });

}

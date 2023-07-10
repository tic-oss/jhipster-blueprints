import fs from 'fs';

let communications = [];

export function loadCommunicationConfigs() {
  const path = this.destinationPath('../' + this.jhipsterConfig.baseName);
  if (this.fs.exists(`${path}/comm.yo-rc.json`)) {
    try {
      const allCommunication = JSON.parse(fs.readFileSync(`${path}/comm.yo-rc.json`));
      communications = allCommunication;
    } catch (err) {
      throw new Error(`Cannot parse the file comm.yo-rc.json in '${this.directoryPath}'`);
    }
  }
}

export function findConfigByBaseName(baseName) {
  const result = [];
  for (const communication of communications) {
    if (communication.client === baseName) {
      result.push({
        server: communication.server,
        type: communication.type,
        framework: communication.framework
      });
    } else if (communication.server === baseName) {
      result.push({
        client: communication.client,
        type: communication.type,
        framework: communication.framework
      });
    }
  }

  return result;
}

export { communications };

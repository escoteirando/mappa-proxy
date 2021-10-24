const fs = require('fs');

const versionFile = 'src/assets/version.js';

/*
version.js

export version = {major:0,minor:0,release:0}
*/
function readVersion() {
  let version = { major: 0, minor: 0, feature: 0 };
  try {
    const data = fs.readFileSync(versionFile, 'utf8');
    const regex = /major: (\d{1,3}), minor: (\d{1,3}), feature: (\d{1,3})/gm;

    let m;
    let matched = 0;
    while ((m = regex.exec(data)) !== null) {
      // This is necessary to avoid infinite loops with zero-width matches
      if (m.index === regex.lastIndex) {
        regex.lastIndex++;
      }

      // The result can be accessed through the `m`-variable.
      m.forEach((match, groupIndex) => {
        switch (groupIndex) {
          case 1:
            version.major = parseInt(match);
            break;
          case 2:
            version.minor = parseInt(match);
            break;
          case 3:
            version.feature = parseInt(match);
            break;
        }
      });
      matched++;
    }
    if (matched == 0) {
      throw Error('Not matched version');
    }
    console.log('Read version', version);
  } catch (err) {
    console.error('Failed to read', { versionFile, err });
  }

  return version;
}

function updateVersion() {
  let version = readVersion();
  if (version.major == 0 && version.minor == 0 && version.feature == 0) {
    console.error('Failed to read version');
  }
  version.feature += 1;
  const updDate = new Date();
  let data = `// Updated - ${updDate.toString()}

export const version = { major: ${version.major}, minor: ${
    version.minor
  }, feature: ${version.feature} };
export const lastUpdate = new Date(${updDate.getTime()})`;
  try {
    fs.writeFileSync(versionFile, data);
    console.info('Version file updated', {
      versionFile,
      version,
      lastUpdate: updDate,
    });
  } catch (err) {
    console.error('Failed to save version file', { versionFile, err });
  }
}

updateVersion();

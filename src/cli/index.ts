import { Arguments } from 'yargs';
import yargs from 'yargs/yargs';
import { hideBin } from 'yargs/helpers';

import logger from '@/services/logger';
import { GithubAPI } from '@/services/api';

import mustache from 'mustache';
import fs from 'graceful-fs';
import _ from 'lodash';

/**
 * Sets the logging level of the `logger` depending on which options
 * where passed
 * @param args The arguments retrieved by yargs
 */
export function setLoggingLevel(args: Arguments) : void {
  if (args.verbose) {
    logger.level = 'verbose';
    logger.warn('Verbose logging has been enabled !');
  } else if (args.silly) {
    logger.level = 'silly';
    logger.warn('Silly logging has been enabled !');
  } else {
    logger.level = 'info';
  }
}

/**
 * Logs a notice when a command is runned
 * @param args The arguments retrieved by `yargs`
 */
export function logExecutionInfo(args: Arguments) : void {
  if (args._.length === 0) {
    return;
  }
  logger.verbose(`Running ${args._[0]} command...`);
  logger.silly('With args: ');
  logger.silly(`\n${JSON.stringify(args, null, 2)}`);
}

/**
 * Initializes a `yargs` based cli with all required configuration
 * @returns The arguments retrieved by `yargs`
 */
export async function initCli(argv: string[] = process.argv) : Promise<Arguments> {
  return yargs(hideBin(argv))
    .scriptName('psmt')
    .version('1.0.0')
    .usage('$0 [command] <options>')
    .demandCommand(0)
    .commandDir('commands')
    .command('login', '', (myargs) => myargs.option('username', {
      type: 'string',
    }), async () => {
      const githubAPI = GithubAPI.create();
      githubAPI.auth();
    })
    .command(
      'template',
      '',
      (myargs) => myargs
        .option('inputPath', {
          type: 'string',
        })
        .option('outputPath', {
          type: 'string',
        })
        .demandOption('inputPath')
        .demandOption('outputPath'),
      async (args) => {
        logger.silly('Starting template command...');
        if (args.inputPath !== undefined) {
          const inputFile = fs.readFileSync(args.inputPath).toString();
          const templateParams = _.chain(mustache.parse(inputFile))
            .map((token) => {
              const name = token[1];
              return name;
            })
            .flatten()
            .uniq()
            .value();
          logger.verbose(`\n${JSON.stringify(templateParams, null, 2)}`);
        }
      },
    )
    .recommendCommands()
    .options({
      verbose: {
        alias: 'v',
        type: 'boolean',
        description: 'Run with verbose level logging',
      },
      silly: {
        type: 'boolean',
        description: 'Run with silly level logging',
        conflicts: 'verbose',
      },
    })
    .help()
    .middleware([
      setLoggingLevel,
      logExecutionInfo,
    ])
    .parse();
}

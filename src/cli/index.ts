import { Arguments } from 'yargs';
import yargs from 'yargs/yargs';
import { hideBin } from 'yargs/helpers';

import logger from '@/services/logger';
import { GithubAPI } from '@/services/api';

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
  logger.verbose(`Running "${args._[0]}" command with args: `);
  logger.verbose(`\n${JSON.stringify(args, null, 2)}`);
}

/**
 * Initializes a `yargs` based cli with all required configuration
 * @returns The arguments retrieved by `yargs`
 */
export async function initCli() : Promise<Arguments> {
  return yargs(hideBin(process.argv))
    .scriptName('psmt')
    .version('1.0.0')
    .usage('$0 [command] <options>')
    .demandCommand(0)
    .commandDir('commands')
    .command('login', '', (myargs) => myargs.option('username', {
      type: 'string',
    }), async (args) => {
      const githubAPI = GithubAPI.create();
      githubAPI.auth();
    })
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

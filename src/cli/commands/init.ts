import logger from '@/services/logger';

export const command = 'init';
export const desc = 'Initialize a psmt config file';
export const builder = '.';

// eslint-disable-next-line no-unused-vars
export const handler = (_args: any) => {
  logger.info('he');
  logger.silly('Yupi');
};

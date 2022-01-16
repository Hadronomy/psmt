import winston from 'winston';
import { merge } from '@/utils';

function format(info : winston.Logform.TransformableInfo) : string {
  return `${info.label || ''}${info.level}: ${info.message}`;
}

export const defaultOptions : winston.LoggerOptions = {
  level: 'info',
  format: winston.format.combine(
    winston.format.timestamp({ format: 'HH:mm:ss.SSS' }),
    winston.format.prettyPrint(),
    winston.format.colorize(),
    winston.format.printf(format),
  ),
  defaultMeta: { service: 'dotfiles' },
  transports: [
    new winston.transports.Console(),
  ],
};

export const defaultLogger : winston.Logger = winston.createLogger(defaultOptions);

export const currentLogger : winston.Logger = defaultLogger;

export function createLogger(options : winston.LoggerOptions) {
  return winston.createLogger(merge(defaultOptions, options));
}

export default currentLogger;

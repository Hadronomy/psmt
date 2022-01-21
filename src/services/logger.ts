import winston from 'winston';
import { merge } from '@/utils';

/**
 * Returns the **default** `logger format`
 * @param info
 * @returns string
 */
function format(info : winston.Logform.TransformableInfo) : string {
  return `${info.label || ''}${info.level} ${info.message}`;
}

/**
 * The **default** `logger options`
 */
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

/**
 * The `logger` that will be used by **default**
 */
export const defaultLogger : winston.Logger = winston.createLogger(defaultOptions);

/**
 * The **currently active** `logger`
 */
export const currentLogger : winston.Logger = defaultLogger;

/**
 * Creates a new `logger` based on the `default options` and the pased `options`
 * provided
 * @param options
 * @returns Logger
 */
export function createLogger(options : winston.LoggerOptions) : winston.Logger {
  return winston.createLogger(merge(defaultOptions, options));
}

export default currentLogger;

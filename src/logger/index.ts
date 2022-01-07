import winston from 'winston';

function format(info : winston.Logform.TransformableInfo) : string {
  return `${info.timestamp} ${info.label || '-'} ${info.level}: ${info.message}`;
}

export const options : winston.LoggerOptions = {
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
export const logger : winston.Logger = winston.createLogger(options);

export default logger;

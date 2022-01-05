import winston from 'winston';

function format(info : winston.Logform.TransformableInfo) : string {
  return `${info.timestamp} ${info.label || '-'} ${info.level}: ${info.message}`;
}

const options : winston.LoggerOptions = {
  level: 'info',
  format: winston.format.combine(
    winston.format.label({ label: '?' }),
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
const logger : winston.Logger = winston.createLogger(options);

export default logger;

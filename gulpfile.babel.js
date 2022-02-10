import {
  series, src, dest,
} from 'gulp';
import fs from 'graceful-fs';
import { posix as path } from 'path';
import babel from 'gulp-babel';
import newer from 'gulp-newer';

const DIST_PATH = path.resolve(__dirname, 'build', 'dist');
const RELEASE_PATH = path.resolve(__dirname, 'build', 'release');
const BABEL_CONF = JSON.parse(fs.readFileSync(path.resolve(__dirname, 'babel.config.json')));

export function clean(cb) {
  if (fs.existsSync(DIST_PATH)) {
    fs.rmSync(DIST_PATH, { recursive: true, force: true });
    fs.mkdirSync(DIST_PATH);
  }
  if (fs.existsSync(RELEASE_PATH)) {
    fs.rmSync(RELEASE_PATH, { recursive: true, force: true });
    fs.mkdirSync(RELEASE_PATH);
  }
  cb();
}

export function build(cb) {
  src('src/**/*')
    .pipe(babel(BABEL_CONF))
    .pipe(newer(DIST_PATH))
    .pipe(dest(DIST_PATH));
  cb();
}

export default series(clean, build);

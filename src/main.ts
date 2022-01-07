#!/bin/node
import { InitCli } from '@/cli';

/**
 * `psmt` main program
 */
async function MainProgram() {
  await InitCli();
}

MainProgram();

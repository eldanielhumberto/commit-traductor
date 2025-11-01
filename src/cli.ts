#!/usr/bin/env bun
import { processCommitMessage } from '.';

const file = process.argv[2];

if (!file) {
	console.error('❌ No se proporcionó el archivo del mensaje de commit.');
	process.exit(1);
}

processCommitMessage(file);

import fs from 'fs';
import { renameCommit } from './renameCommit';

export const processCommitMessage = async (file: string) => {
	const msg = fs.readFileSync(file, 'utf8').trim();
	const newMsg = await renameCommit(msg);

	fs.writeFileSync(file, newMsg, 'utf8');
	console.log(`Commit translated to: ${newMsg}\n`);
};

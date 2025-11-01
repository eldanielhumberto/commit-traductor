import fs from 'fs';
import translate from 'translate';

export async function renameCommit(msg: string): Promise<string> {
	translate.engine = 'google';
	return await translate(msg, { from: 'es', to: 'en' });
}

export async function processCommitMessage(file: string) {
	const msg = fs.readFileSync(file, 'utf8').trim();
	const newMsg = await renameCommit(msg);

	fs.writeFileSync(file, newMsg, 'utf8');
	console.log(`✅ Commit translated to: ${newMsg}\n`);
}

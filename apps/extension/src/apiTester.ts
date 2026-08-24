import * as vscode from 'vscode';
import { APIContractItem } from './apiExplorer';

export function registerAPITester(context: vscode.ExtensionContext) {
    const disposable = vscode.commands.registerCommand('devhub.testEndpoint', async (item?: APIContractItem) => {
        const endpointPath = item?.path || '/api/v1/tasks';
        const method = item?.method || 'GET';

        try {
            const url = `http://localhost:8080${endpointPath}`;
            const res = await fetch(url);
            const json = await res.json();

            const doc = await vscode.workspace.openTextDocument({
                content: JSON.stringify({
                    request: { method, url },
                    status: res.status,
                    headers: Object.fromEntries(res.headers.entries()),
                    body: json
                }, null, 2),
                language: 'json'
            });

            await vscode.window.showTextDocument(doc, { preview: true });
        } catch (err: any) {
            vscode.window.showErrorMessage(`API Test Failed: ${err.message}`);
        }
    });

    context.subscriptions.push(disposable);
}

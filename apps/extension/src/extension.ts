import * as vscode from 'vscode';

export function activate(context: vscode.ExtensionContext) {
    console.log('devHub Extension is now active!');

    let disposable = vscode.commands.registerCommand('devhub.helloWorld', () => {
        vscode.window.showInformationMessage('Hello from devHub Internal Developer Portal!');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}

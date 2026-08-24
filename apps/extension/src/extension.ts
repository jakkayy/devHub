import * as vscode from 'vscode';
import { DevHubSidebarProvider } from './sidebarProvider';

export function activate(context: vscode.ExtensionContext) {
    console.log('devHub Extension is now active!');

    const sidebarProvider = new DevHubSidebarProvider(context.extensionUri);
    
    context.subscriptions.push(
        vscode.window.registerWebviewViewProvider(
            DevHubSidebarProvider.viewType,
            sidebarProvider
        )
    );

    let disposable = vscode.commands.registerCommand('devhub.helloWorld', () => {
        vscode.window.showInformationMessage('Hello from devHub Internal Developer Portal!');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}

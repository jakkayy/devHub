import * as vscode from 'vscode';
import { DevHubSidebarProvider } from './sidebarProvider';
import { APIContractTreeDataProvider } from './apiExplorer';
import { LocalMockServer } from './mockServer';
import { registerAPITester } from './apiTester';
import { registerTaskLinkerCommand } from './taskLinker';
import { DevHubCodeLensProvider } from './codeLensProvider';
import { registerContextDetector } from './contextDetector';
import { registerNotificationTrigger } from './notifier';

export function activate(context: vscode.ExtensionContext) {
    console.log('devHub Extension is now active!');

    const sidebarProvider = new DevHubSidebarProvider(context.extensionUri);
    const apiExplorerProvider = new APIContractTreeDataProvider();
    const mockServer = new LocalMockServer();
    const codeLensProvider = new DevHubCodeLensProvider();

    mockServer.startServer();
    registerAPITester(context);
    registerTaskLinkerCommand(context);
    registerContextDetector(context);
    registerNotificationTrigger(context);

    context.subscriptions.push(
        vscode.languages.registerCodeLensProvider({ scheme: 'file' }, codeLensProvider),
        vscode.window.registerWebviewViewProvider(
            DevHubSidebarProvider.viewType,
            sidebarProvider
        ),
        vscode.window.registerTreeDataProvider(
            'devhub.apiExplorerView',
            apiExplorerProvider
        ),
        { dispose: () => mockServer.stopServer() }
    );

    let disposable = vscode.commands.registerCommand('devhub.startMockServer', () => {
        vscode.window.showInformationMessage('devHub Local OpenAPI Mock Server running on http://localhost:9090');
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}

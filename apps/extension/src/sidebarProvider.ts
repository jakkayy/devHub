import * as vscode from 'vscode';

export class DevHubSidebarProvider implements vscode.WebviewViewProvider {
    public static readonly viewType = 'devhub.sidebarView';
    private _view?: vscode.WebviewView;

    constructor(private readonly _extensionUri: vscode.Uri) {}

    public resolveWebviewView(
        webviewView: vscode.WebviewView,
        context: vscode.WebviewViewResolveContext,
        _token: vscode.CancellationToken
    ) {
        this._view = webviewView;

        webviewView.webview.options = {
            enableScripts: true,
            localResourceRoots: [this._extensionUri]
        };

        webviewView.webview.html = this._getHtmlForWebview(webviewView.webview);
    }

    public postMessageToWebview(message: any) {
        if (this._view) {
            this._view.webview.postMessage(message);
        }
    }

    private _getHtmlForWebview(webview: vscode.Webview): string {
        return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>devHub Command Center</title>
    <style>
        body {
            font-family: var(--vscode-font-family);
            padding: 10px;
            color: var(--vscode-foreground);
            background-color: var(--vscode-sideBar-background);
        }
        .header {
            font-weight: bold;
            margin-bottom: 12px;
            font-size: 1.1em;
            display: flex;
            align-items: center;
            gap: 6px;
        }
        .status-badge {
            font-size: 0.8em;
            padding: 2px 6px;
            border-radius: 10px;
            background: rgba(16, 185, 129, 0.2);
            color: #10b981;
            border: 1px solid rgba(16, 185, 129, 0.4);
        }
    </style>
</head>
<body>
    <div class="header">
        <span>devHub Sidebar</span>
        <span class="status-badge">Active</span>
    </div>
    <div id="app">Initializing devHub Command Center...</div>
</body>
</html>`;
    }
}

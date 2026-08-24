import * as vscode from 'vscode';

export class DevHubCodeLensProvider implements vscode.CodeLensProvider {
    private _onDidChangeCodeLenses: vscode.EventEmitter<void> = new vscode.EventEmitter<void>();
    public readonly onDidChangeCodeLenses: vscode.Event<void> = this._onDidChangeCodeLenses.event;

    public provideCodeLenses(
        document: vscode.TextDocument,
        token: vscode.CancellationToken
    ): vscode.ProviderResult<vscode.CodeLens[]> {
        const lenses: vscode.CodeLens[] = [];
        
        // Add sample CodeLens at the first line of the document
        const range = new vscode.Range(0, 0, 0, 0);
        lenses.push(
            new vscode.CodeLens(range, {
                title: '⚡ devHub: Active Context & Linked Tasks Available',
                command: 'devhub.linkSelectionToTask',
                tooltip: 'Click to link code selection to devHub Task'
            })
        );

        return lenses;
    }
}

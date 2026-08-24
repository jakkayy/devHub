import * as vscode from 'vscode';

export function registerContextDetector(context: vscode.ExtensionContext) {
    const disposable = vscode.window.onDidChangeActiveTextEditor((editor) => {
        if (!editor) {
            return;
        }

        const fileName = editor.document.fileName.split('/').pop();
        console.log(`devHub Context Detector: Active Editor changed to ${fileName}`);
        
        // Show subtle status bar message for context awareness
        vscode.window.setStatusBarMessage(`devHub Context: ${fileName}`, 3000);
    });

    context.subscriptions.push(disposable);
}

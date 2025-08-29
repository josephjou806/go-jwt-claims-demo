# AI Code Review Configuration

This directory contains the configuration and scripts for automated code review using Claude AI.

## Files

- `claude-review.js` - Main script that performs the AI code review
- `package.json` - Node.js dependencies for the review script
- `review-config.json` - Configuration for review parameters

## Setup Instructions

### 1. GitHub Repository Secrets

Add the following secrets to your GitHub repository (Settings → Secrets and variables → Actions):

- `ANTHROPIC_API_KEY` - Your Anthropic API key for Claude access

### 2. GitHub Token Permissions

The workflow uses the default `GITHUB_TOKEN` with the following permissions:
- `contents: read` - To access repository files
- `pull-requests: write` - To post review comments
- `issues: write` - To create issue comments

### 3. Anthropic API Key

1. Go to [Anthropic Console](https://console.anthropic.com/)
2. Create an account or sign in
3. Generate an API key
4. Add it to your GitHub repository secrets as `ANTHROPIC_API_KEY`

## How It Works

1. **Trigger**: The workflow runs on every pull request (opened, synchronized, reopened)
2. **Analysis**: Changed files are analyzed by Claude AI focusing on:
   - Security vulnerabilities
   - Performance optimizations
   - Code maintainability
   - Go idioms and conventions
   - Error handling patterns
   - Testing coverage
   - Documentation completeness

3. **Review**: Claude provides detailed feedback including:
   - Issues found with severity levels
   - Security considerations
   - Performance suggestions
   - Testing recommendations
   - Overall recommendation

4. **Comment**: The review is posted as a comment on the pull request

## Customization

You can customize the review behavior by modifying the `REVIEW_CONFIG` object in `claude-review.js`:

```javascript
const REVIEW_CONFIG = {
  maxFilesPerReview: 10,           // Maximum files to review per PR
  maxCharsPerFile: 8000,           // Maximum characters per file
  focusAreas: [...],               // Areas to focus on during review
  excludePatterns: [...]           // File patterns to exclude
};
```

## Local Testing

To test the review script locally:

```bash
cd .github/scripts
npm install
export ANTHROPIC_API_KEY="your-api-key"
export GITHUB_TOKEN="your-github-token"
export PR_NUMBER="123"
export REPO_OWNER="your-username"
export REPO_NAME="your-repo"
node claude-review.js
```

## Troubleshooting

### Common Issues

1. **API Key Issues**
   - Ensure `ANTHROPIC_API_KEY` is correctly set in repository secrets
   - Verify the API key has sufficient credits

2. **Permission Issues**
   - Check that the workflow has the required permissions
   - Ensure the repository allows workflows to post comments

3. **Review Quality**
   - Adjust `maxCharsPerFile` if reviews are truncated
   - Modify `focusAreas` to target specific concerns
   - Update the prompt in `claude-review.js` for different review styles

### Workflow Logs

Check the GitHub Actions logs for detailed error information:
1. Go to your repository
2. Click "Actions" tab
3. Select the failed workflow run
4. Examine the "Run AI Code Review" step logs

## Best Practices

1. **Use as Supplement**: AI reviews should complement, not replace, human code reviews
2. **Regular Updates**: Keep the Claude model version and dependencies updated
3. **Security**: Never commit API keys to the repository
4. **Costs**: Monitor Anthropic API usage to manage costs
5. **Feedback Loop**: Regularly review and improve the prompts based on review quality

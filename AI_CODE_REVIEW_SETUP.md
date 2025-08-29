# Go JWT Claims Demo - AI Code Review Setup

This repository now includes an automated AI code review process using Claude AI and GitHub workflows.

## 🚀 What's Included

### 1. GitHub Workflow (`.github/workflows/code-review.yml`)
- Triggers on pull requests to `main` and `develop` branches
- Runs Go tests, linting, and static analysis
- Performs AI-powered code review using Claude
- Posts detailed review comments on pull requests

### 2. Claude AI Review Script (`.github/scripts/claude-review.js`)
- Node.js script that interfaces with Anthropic's Claude API
- Analyzes code changes with focus on Go best practices
- Provides detailed feedback on security, performance, and maintainability
- Customizable review criteria and severity levels

### 3. Configuration Files
- `package.json` - Dependencies for the review script
- `review-config.json` - Customizable review parameters
- `README.md` - Detailed setup and usage instructions

## 🛠️ Setup Instructions

### Step 1: Get Anthropic API Key
1. Visit [Anthropic Console](https://console.anthropic.com/)
2. Create an account or sign in
3. Generate an API key
4. Note down the API key (you'll need it in Step 2)

### Step 2: Configure GitHub Repository
1. Go to your repository on GitHub
2. Navigate to **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add the following secret:
   - Name: `ANTHROPIC_API_KEY`
   - Value: Your Anthropic API key from Step 1

### Step 3: Enable Workflow Permissions
1. In your repository, go to **Settings** → **Actions** → **General**
2. Under "Workflow permissions", ensure:
   - "Read and write permissions" is selected, OR
   - "Read repository contents and packages permissions" with additional permissions for:
     - Contents: Read
     - Issues: Write  
     - Pull requests: Write

### Step 4: Test the Setup
1. Create a test branch:
   ```bash
   git checkout -b test-ai-review
   ```
2. Make a small change to any `.go` file
3. Commit and push:
   ```bash
   git add .
   git commit -m "Test AI code review setup"
   git push origin test-ai-review
   ```
4. Create a pull request from `test-ai-review` to `main`
5. The workflow should automatically trigger and post a review comment

## 🎯 What the AI Reviews

The Claude AI reviewer focuses on:

### Security 🔒
- JWT token handling and validation
- Input sanitization and validation
- CORS configuration
- Secret management
- Authentication/authorization patterns

### Performance 🚀
- Memory usage optimization
- Goroutine management
- Database query efficiency
- Caching strategies
- Resource cleanup

### Code Quality 📝
- Go idioms and conventions
- Error handling patterns
- Testing coverage and quality
- Documentation completeness
- API design consistency

### Best Practices ✅
- Dependency management
- Code maintainability
- Concurrency safety
- Project structure
- Configuration management

## 🔧 Customization

### Modify Review Focus
Edit `.github/scripts/review-config.json` to customize:
- Maximum files per review
- File exclusion patterns
- Focus areas and priorities
- Severity level criteria

### Adjust AI Behavior
In `.github/scripts/claude-review.js`, you can:
- Modify the Claude model and parameters
- Customize the review prompt
- Adjust output format
- Add project-specific context

### Example Custom Focus Areas
```json
{
  "focusAreas": [
    "JWT security best practices",
    "Gin framework optimization",
    "Database connection handling",
    "Error response consistency",
    "API versioning strategy"
  ]
}
```

## 📊 Review Output Example

The AI will post comments like this on your pull requests:

```markdown
## 🤖 AI Code Review by Claude

### ✅ Positive Aspects
- Well-structured JWT middleware implementation
- Proper error handling in authentication flow
- Good separation of concerns between handlers and services

### ⚠️ Issues Found
- **[Severity: High]** JWT secret should not be hardcoded
  - File: `internal/config/config.go`
  - Suggestion: Use environment variables for sensitive configuration

### 🔒 Security Considerations
- Consider adding JWT token expiration validation
- Implement rate limiting for authentication endpoints

### 🎯 Overall Recommendation
- [x] Request changes (specific issues must be addressed)
```

## 💡 Best Practices

1. **Use as Supplement**: AI reviews complement but don't replace human reviews
2. **Monitor Costs**: Track your Anthropic API usage
3. **Iterate and Improve**: Refine prompts based on review quality
4. **Security First**: Never commit API keys to your repository
5. **Team Alignment**: Ensure team understands AI review limitations

## 🐛 Troubleshooting

### Common Issues
- **API Key Error**: Verify `ANTHROPIC_API_KEY` is correctly set in repository secrets
- **Permission Denied**: Check workflow permissions in repository settings
- **No Review Posted**: Ensure the PR contains changes to reviewable files (.go, .md, etc.)

### Getting Help
- Check GitHub Actions logs for detailed error information
- Review the setup documentation in `.github/scripts/README.md`
- Verify API key has sufficient credits at Anthropic Console

---

Ready to experience AI-powered code reviews! 🎉

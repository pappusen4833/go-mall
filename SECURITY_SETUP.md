# Security Configuration Guide / 安全配置指南

## ⚠️ Important Security Notice / 重要安全提示

This project requires proper configuration of sensitive credentials before running. **Never commit real credentials to the repository.**

本项目需要正确配置敏感凭据才能运行。**切勿将真实凭据提交到代码仓库。**

## Configuration Setup / 配置设置

### 1. Create Configuration File / 创建配置文件

Copy the example configuration file and update it with your credentials:

复制示例配置文件并使用您的凭据更新它:

```bash
cp config.example.yaml config.yaml
```

### 2. Update Configuration / 更新配置

Edit `config.yaml` and replace all placeholder values with your actual credentials:

编辑 `config.yaml` 并将所有占位符替换为您的实际凭据:

#### Database Configuration / 数据库配置
```yaml
database:
  user: 'your_actual_database_user'
  password: 'your_actual_database_password'
  host: '127.0.0.1:3306'
  name: 'gomall'
```

#### WeChat Configuration / 微信配置
```yaml
wechat:
  app_id: your_actual_wechat_app_id
  app_secret: your_actual_wechat_app_secret
  token: your_actual_wechat_token
```

#### WeChat Pay Configuration / 微信支付配置
```yaml
wxpay:
  mch_id: your_actual_merchant_id
  api_key: your_actual_api_key
  notify_url: https://your-domain.com/wxPay/notify
```

#### Express API Configuration / 快递API配置
```yaml
express:
  eBusinessId: your_actual_express_business_id
  appKey: your_actual_express_app_key
```

#### JWT Secret / JWT密钥
```yaml
app:
  jwt-secret: generate_a_strong_random_secret_here
```

**Tip**: Use a strong random string for jwt-secret (at least 32 characters)

**提示**: 为 jwt-secret 使用强随机字符串(至少32个字符)

### 3. Security Checklist / 安全检查清单

Before deploying or committing code, ensure:

在部署或提交代码之前,确保:

- [x] `config.yaml` is listed in `.gitignore`
- [x] No hardcoded credentials in source code
- [x] All sensitive values use configuration file
- [x] Strong passwords and secrets are used
- [x] Production credentials differ from development

### 4. Environment Variables (Optional) / 环境变量(可选)

You can also use environment variables to override configuration:

您也可以使用环境变量覆盖配置:

```bash
export DB_PASSWORD=your_password
export WECHAT_APP_SECRET=your_secret
export JWT_SECRET=your_jwt_secret
```

### 5. Database Initialization / 数据库初始化

⚠️ **Important**: The provided SQL file (`sql/go-mall.sql`) contains test data including:

重要提示: 提供的SQL文件 (`sql/go-mall.sql`) 包含测试数据:

- Test user accounts with hashed passwords
- Sample email addresses and phone numbers
- Demo product and order data

**Recommendations / 建议:**

1. **For Development**: Use the provided SQL file as-is
2. **For Production**:
   - Create your own clean database schema
   - Remove or anonymize test user data
   - Update admin credentials after initial setup

**Production Setup / 生产环境设置:**

```sql
-- After importing the SQL file, update admin password
UPDATE sys_user SET password = 'your_new_bcrypt_hashed_password' WHERE username = 'admin';

-- Remove test users if not needed
DELETE FROM sys_user WHERE id > 1;
```

### 6. Git History Cleanup / Git历史清理

If you accidentally committed sensitive data, clean it from git history:

如果您不小心提交了敏感数据,请从git历史记录中清除它:

```bash
# Remove config.yaml from all commits
git filter-branch --force --index-filter \
  'git rm --cached --ignore-unmatch config.yaml' \
  --prune-empty --tag-name-filter cat -- --all

# Force push (use with caution!)
git push origin --force --all
```

## Best Practices / 最佳实践

1. **Never commit `config.yaml`** - Always use `config.example.yaml` as template
2. **Use strong passwords** - Minimum 16 characters with mixed case, numbers, symbols
3. **Rotate credentials regularly** - Change passwords and API keys periodically
4. **Separate environments** - Use different credentials for dev/staging/production
5. **Use environment variables in production** - Avoid storing secrets in files
6. **Enable 2FA** - Enable two-factor authentication for critical services
7. **Monitor access logs** - Regularly review access logs for suspicious activity

## Getting Help / 获取帮助

If you discover a security vulnerability, please email security@your-domain.com instead of creating a public issue.

如果您发现安全漏洞,请发送电子邮件至 security@your-domain.com 而不是创建公开问题。

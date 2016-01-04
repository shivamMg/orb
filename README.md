# orb

Slack bot to post CodeChef programming problems at 8:30 each day. **orb** uses cron job for scheduling.

## Openshift Deployment

**orb** can be deployed on [Openshift](https://openshift.com) with the following steps. Make sure you have [Openshift Client Tools](https://developers.openshift.com/en/managing-client-tools.html) installed.

1. Install the Go cartridge.

   ```
   rhc create-app orb https://cartreflect-claytondev.rhcloud.com/reflect?github=smarterclayton/openshift-go-cart
   ```

2. Create a local copy of your project.

   ```
   rhc git-clone -a orb
   ```

3. Update `.godir` file with your package name. e.g. The `.godir` file might contain:

   ```
   github.com/shivamMg/orb
   ```

4. Install cron cartridge for your application.

   ```
   rhc cartridge add cron -a orb
   ```

5. You need to add a cron job that runs hourly, so create `hourly` directory inside `cron` and create a bash script in it.

   ```
   mkdir -p .openshift/cron/hourly
   touch .openshift/cron/hourly/script.sh
   ```

6. Add the following inside the script:

   ```bash
   #! /bin/bash

   export TZ="Asia/Calcutta"
   date >> $OPENSHIFT_GO_LOG_DIR/last_date_cron_ran
   $OPENSHIFT_REPO_DIR/bin/orb >> $OPENSHIFT_GO_LOG_DIR/last_date_cron_ran
   ```

   This will keep a track of the questions posted.

7. Make your script executable.

   ```
   chmod +x .openshift/cron/hourly/script.sh
   ```

8. Add necessary files used by **orb**.

9. Add your changes, commit work and push.

   ```
   git add *
   git commit -m "Add orb"
   git push
   ```

## License

See `LICENSE` file.


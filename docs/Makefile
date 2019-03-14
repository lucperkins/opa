setup:
	yarn global add postcss-cli

serve:
	hugo server \
		--buildDrafts \
		--buildFuture \
		--disableFastRender

production-build: setup
	hugo --minify

preview-build: setup
	hugo \
		--baseURL $(DEPLOY_PRIME_URL) \
		--buildDrafts \
		--buildFuture \
		--minify

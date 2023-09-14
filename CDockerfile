# Use the official Chrome headless image as a parent image
FROM zenika/alpine-chrome:latest

# Expose the port Chrome DevTools Protocol runs on
EXPOSE 9222

# Set a label to disable the sandbox, which is necessary for Chrome to run in a container
LABEL chromeoptions.chromium-browser.sandbox.enabled="false"

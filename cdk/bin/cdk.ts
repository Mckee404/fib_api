#!/usr/bin/env node
import * as cdk from 'aws-cdk-lib';
import { FibApiStack } from '../lib/cdk-stack';

const app = new cdk.App();
const stage = app.node.tryGetContext('stage') || 'dev';

new FibApiStack(app, `FibApiStack-${stage}`, {
  stage: stage,
  env: {
    region: 'ap-northeast-1',
  },
});
# Bosh-cpi-cloudstack

Complete rewrite in golang of [current cpi cloudstack](https://github.com/cloudfoundry-community/bosh-cloudstack-cpi-core/).

This new CPI remove boilerplate as:
- **Webdav**: This was used to be able to register stemcell as template in cloudstack.
We now use template upload provided natively by cloudstack (see: https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=39620237)
- **Registry**: There is no [registry](https://bosh.io/docs/bosh-components/#registry) implementation in this CPI.
We target to use registry provided by bosh.
- **Web Server**: This new CPI is not anymore a webserver as the previous CPI. This is fully compliant with bosh RPC mechanism to call CPI.

## Install

Please use boshrelease directly available at: https://github.com/orange-cloudfoundry/bosh-go-cpi-cloudstack-cpi
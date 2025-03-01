// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ProviderCustomization {
    setAddonConfigVarsInState?: pulumi.Input<boolean>;
    setAppAllConfigVarsInState?: pulumi.Input<boolean>;
}

export interface ProviderDelays {
    postAppCreateDelay?: pulumi.Input<number>;
    postDomainCreateDelay?: pulumi.Input<number>;
    postSpaceCreateDelay?: pulumi.Input<number>;
}

export interface ProviderTimeouts {
    addonCreateTimeout?: pulumi.Input<number>;
}
export namespace app {
    export interface AppOrganization {
        /**
         * Are other team members forbidden from joining this app.
         */
        locked?: pulumi.Input<boolean>;
        /**
         * The name of the Heroku Team.
         */
        name: pulumi.Input<string>;
        /**
         * Force creation of the app in the user account even if a default team is set.
         */
        personal?: pulumi.Input<boolean>;
    }

}

export namespace build {
    export interface BuildSource {
        /**
         * SHA256 hash of the tarball archive to verify its integrity, example:
         * `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
         */
        checksum?: pulumi.Input<string>;
        /**
         * Local path to the source directory or tarball archive for the app
         */
        path?: pulumi.Input<string>;
        /**
         * `https` location of the source archive for the app
         */
        url?: pulumi.Input<string>;
        /**
         * Use to track what version of your source originated this build. If you are creating builds
         * from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
         * version parameter.
         */
        version?: pulumi.Input<string>;
    }

    export interface BuildUser {
        email?: pulumi.Input<string>;
        id?: pulumi.Input<string>;
    }
}

export namespace config {
}

export namespace pipeline {
    export interface PipelineOwner {
        /**
         * The unique identifier (UUID) of a pipeline owner.
         */
        id: pulumi.Input<string>;
        /**
         * The type of pipeline owner. Can be either `user` or `team`.
         */
        type: pulumi.Input<string>;
    }
}

export namespace review {
    export interface AppConfigDeployTarget {
        /**
         * Unique identifier of deploy target.
         */
        id: pulumi.Input<string>;
        /**
         * Type of deploy target. Must be either `space` or `region`.
         */
        type: pulumi.Input<string>;
    }
}

export namespace slug {
    export interface SlugBlob {
        /**
         * HTTP method to upload the archive
         */
        method?: pulumi.Input<string>;
        /**
         * Pre-signed, expiring URL to upload the archive
         */
        url?: pulumi.Input<string>;
    }
}

export namespace space {
    export interface InboundRulesetRule {
        /**
         * The action to apply this rule to. Must be one of `allow` or `deny`.
         */
        action: pulumi.Input<string>;
        /**
         * A CIDR block source for the rule.
         */
        source: pulumi.Input<string>;
    }

    export interface VpnConnectionTunnel {
        /**
         * The public IP address of the tunnel.
         */
        ip?: pulumi.Input<string>;
        /**
         * The pre-shared IPSec secret for the tunnel.
         */
        preSharedKey?: pulumi.Input<string>;
    }
}

export namespace team {
}

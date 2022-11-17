// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace app {
    export interface AppOrganization {
        locked: boolean;
        name: string;
        personal: boolean;
    }

    export interface GetAppOrganization {
        locked: boolean;
        name: string;
        personal: boolean;
    }

}

export namespace build {
    export interface BuildSource {
        checksum: string;
        path?: string;
        url?: string;
        version?: string;
    }

    export interface BuildUser {
        email: string;
        id: string;
    }

}

export namespace config {
    export interface Customizations {
        setAppAllConfigVarsInState?: boolean;
    }

    export interface Delays {
        postAppCreateDelay?: number;
        postDomainCreateDelay?: number;
        postSpaceCreateDelay?: number;
    }

    export interface Timeouts {
        addonCreateTimeout?: number;
    }

}

export namespace pipeline {
    export interface PipelineOwner {
        id: string;
        type: string;
    }

}

export namespace review {
    export interface AppConfigDeployTarget {
        id: string;
        type: string;
    }

}

export namespace slug {
    export interface SlugBlob {
        method: string;
        url: string;
    }

}

export namespace space {
    export interface InboundRulesetRule {
        action: string;
        source: string;
    }

    export interface VpnConnectionTunnel {
        ip: string;
        preSharedKey: string;
    }

}

export namespace team {
    export interface GetMembersMember {
        email: string;
        federated: boolean;
        role: string;
        teamMemberId: string;
        twoFactorAuthentication: boolean;
        userId: string;
    }

}

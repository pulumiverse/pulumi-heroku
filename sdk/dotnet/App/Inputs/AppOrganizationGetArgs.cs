// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.App.Inputs
{

    public sealed class AppOrganizationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Are other team members forbidden from joining this app.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        /// <summary>
        /// The name of the Heroku Team.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Force creation of the app in the user account even if a default team is set.
        /// </summary>
        [Input("personal")]
        public Input<bool>? Personal { get; set; }

        public AppOrganizationGetArgs()
        {
        }
        public static new AppOrganizationGetArgs Empty => new AppOrganizationGetArgs();
    }
}

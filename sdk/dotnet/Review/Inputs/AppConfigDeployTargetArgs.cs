// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Review.Inputs
{

    public sealed class AppConfigDeployTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of deploy target.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Type of deploy target. Must be either `space` or `region`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AppConfigDeployTargetArgs()
        {
        }
        public static new AppConfigDeployTargetArgs Empty => new AppConfigDeployTargetArgs();
    }
}

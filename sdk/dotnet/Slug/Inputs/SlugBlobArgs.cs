// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Slug.Inputs
{

    public sealed class SlugBlobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP method to upload the archive
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// Pre-signed, expiring URL to upload the archive
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public SlugBlobArgs()
        {
        }
        public static new SlugBlobArgs Empty => new SlugBlobArgs();
    }
}

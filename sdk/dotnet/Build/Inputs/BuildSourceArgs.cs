// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Build.Inputs
{

    public sealed class BuildSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SHA256 hash of the tarball archive to verify its integrity, example:
        /// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
        /// </summary>
        [Input("checksum")]
        public Input<string>? Checksum { get; set; }

        /// <summary>
        /// Local path to the source directory or tarball archive for the app
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// `https` location of the source archive for the app
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// Use to track what version of your source originated this build. If you are creating builds
        /// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
        /// version parameter.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public BuildSourceArgs()
        {
        }
        public static new BuildSourceArgs Empty => new BuildSourceArgs();
    }
}

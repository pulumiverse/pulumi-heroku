// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Build.Outputs
{

    [OutputType]
    public sealed class BuildSource
    {
        /// <summary>
        /// SHA256 hash of the tarball archive to verify its integrity, example:
        /// `SHA256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`
        /// </summary>
        public readonly string? Checksum;
        /// <summary>
        /// Local path to the source directory or tarball archive for the app
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// `https` location of the source archive for the app
        /// </summary>
        public readonly string? Url;
        /// <summary>
        /// Use to track what version of your source originated this build. If you are creating builds
        /// from git-versioned source code, for example, the commit hash, or release tag would be a good value to use for the
        /// version parameter.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private BuildSource(
            string? checksum,

            string? path,

            string? url,

            string? version)
        {
            Checksum = checksum;
            Path = path;
            Url = url;
            Version = version;
        }
    }
}

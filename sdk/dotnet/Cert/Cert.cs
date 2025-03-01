// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Cert
{
    /// <summary>
    /// This resource manages an SSL certificate for a Heroku app.
    /// 
    /// !&gt; **WARNING:** This resource is deprecated in favor of `heroku.ssl.Ssl`.
    /// 
    /// ## Example Usage
    /// 
    /// ## Importing
    /// 
    /// When importing a Heroku cert resource, the ID must be built using the app name colon the unique ID from the Heroku API. For an app named `production-api` with a certificate ID of `b85d9224-310b-409b-891e-c903f5a40568`, you would import it as:
    /// </summary>
    [HerokuResourceType("heroku:cert/cert:Cert")]
    public partial class Cert : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Output("app")]
        public Output<string> App { get; private set; } = null!;

        /// <summary>
        /// The certificate chain to add
        /// </summary>
        [Output("certificateChain")]
        public Output<string> CertificateChain { get; private set; } = null!;

        /// <summary>
        /// The CNAME for the SSL endpoint
        /// </summary>
        [Output("cname")]
        public Output<string> Cname { get; private set; } = null!;

        /// <summary>
        /// The name of the SSL certificate
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The private key for a given certificate chain
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;


        /// <summary>
        /// Create a Cert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Cert(string name, CertArgs args, CustomResourceOptions? options = null)
            : base("heroku:cert/cert:Cert", name, args ?? new CertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Cert(string name, Input<string> id, CertState? state = null, CustomResourceOptions? options = null)
            : base("heroku:cert/cert:Cert", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "privateKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Cert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Cert Get(string name, Input<string> id, CertState? state = null, CustomResourceOptions? options = null)
        {
            return new Cert(name, id, state, options);
        }
    }

    public sealed class CertArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("app", required: true)]
        public Input<string> App { get; set; } = null!;

        /// <summary>
        /// The certificate chain to add
        /// </summary>
        [Input("certificateChain", required: true)]
        public Input<string> CertificateChain { get; set; } = null!;

        [Input("privateKey", required: true)]
        private Input<string>? _privateKey;

        /// <summary>
        /// The private key for a given certificate chain
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CertArgs()
        {
        }
        public static new CertArgs Empty => new CertArgs();
    }

    public sealed class CertState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Heroku app ID (do not use app name)
        /// </summary>
        [Input("app")]
        public Input<string>? App { get; set; }

        /// <summary>
        /// The certificate chain to add
        /// </summary>
        [Input("certificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// The CNAME for the SSL endpoint
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// The name of the SSL certificate
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateKey")]
        private Input<string>? _privateKey;

        /// <summary>
        /// The private key for a given certificate chain
        /// </summary>
        public Input<string>? PrivateKey
        {
            get => _privateKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _privateKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public CertState()
        {
        }
        public static new CertState Empty => new CertState();
    }
}

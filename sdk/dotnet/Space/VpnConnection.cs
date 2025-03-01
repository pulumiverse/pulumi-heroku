// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Heroku.Space
{
    /// <summary>
    /// Provides a resource for creating a VPN connection between a network and a Heroku Private Space. For more information, see [Private Spaces VPN Connection](https://devcenter.heroku.com/articles/private-space-vpn-connection?preview=1) in the Heroku DevCenter.
    /// 
    /// ## Example Usage
    /// </summary>
    [HerokuResourceType("heroku:space/vpnConnection:VpnConnection")]
    public partial class VpnConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IKE version used to setup the IPsec tunnel.
        /// </summary>
        [Output("ikeVersion")]
        public Output<int> IkeVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The public IP address of the VPN endpoint on the network where the VPN connection will be established.
        /// </summary>
        [Output("publicIp")]
        public Output<string> PublicIp { get; private set; } = null!;

        /// <summary>
        /// A list of IPv4 CIDR blocks used by the network where the VPN connection will be established.
        /// </summary>
        [Output("routableCidrs")]
        public Output<ImmutableArray<string>> RoutableCidrs { get; private set; } = null!;

        /// <summary>
        /// The ID of the Heroku Private Space where the VPN connection will be established.
        /// </summary>
        [Output("space")]
        public Output<string> Space { get; private set; } = null!;

        /// <summary>
        /// The CIDR block for the Heroku Private Space. The network where the VPN will be established should be configured to route traffic destined for this CIDR block over the VPN link.
        /// </summary>
        [Output("spaceCidrBlock")]
        public Output<string> SpaceCidrBlock { get; private set; } = null!;

        /// <summary>
        /// Details about each VPN tunnel endpoint.
        /// </summary>
        [Output("tunnels")]
        public Output<ImmutableArray<Outputs.VpnConnectionTunnel>> Tunnels { get; private set; } = null!;


        /// <summary>
        /// Create a VpnConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnConnection(string name, VpnConnectionArgs args, CustomResourceOptions? options = null)
            : base("heroku:space/vpnConnection:VpnConnection", name, args ?? new VpnConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnConnection(string name, Input<string> id, VpnConnectionState? state = null, CustomResourceOptions? options = null)
            : base("heroku:space/vpnConnection:VpnConnection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnConnection Get(string name, Input<string> id, VpnConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new VpnConnection(name, id, state, options);
        }
    }

    public sealed class VpnConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public IP address of the VPN endpoint on the network where the VPN connection will be established.
        /// </summary>
        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        [Input("routableCidrs", required: true)]
        private InputList<string>? _routableCidrs;

        /// <summary>
        /// A list of IPv4 CIDR blocks used by the network where the VPN connection will be established.
        /// </summary>
        public InputList<string> RoutableCidrs
        {
            get => _routableCidrs ?? (_routableCidrs = new InputList<string>());
            set => _routableCidrs = value;
        }

        /// <summary>
        /// The ID of the Heroku Private Space where the VPN connection will be established.
        /// </summary>
        [Input("space", required: true)]
        public Input<string> Space { get; set; } = null!;

        [Input("tunnels")]
        private InputList<Inputs.VpnConnectionTunnelArgs>? _tunnels;

        /// <summary>
        /// Details about each VPN tunnel endpoint.
        /// </summary>
        public InputList<Inputs.VpnConnectionTunnelArgs> Tunnels
        {
            get => _tunnels ?? (_tunnels = new InputList<Inputs.VpnConnectionTunnelArgs>());
            set => _tunnels = value;
        }

        public VpnConnectionArgs()
        {
        }
        public static new VpnConnectionArgs Empty => new VpnConnectionArgs();
    }

    public sealed class VpnConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IKE version used to setup the IPsec tunnel.
        /// </summary>
        [Input("ikeVersion")]
        public Input<int>? IkeVersion { get; set; }

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public IP address of the VPN endpoint on the network where the VPN connection will be established.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        [Input("routableCidrs")]
        private InputList<string>? _routableCidrs;

        /// <summary>
        /// A list of IPv4 CIDR blocks used by the network where the VPN connection will be established.
        /// </summary>
        public InputList<string> RoutableCidrs
        {
            get => _routableCidrs ?? (_routableCidrs = new InputList<string>());
            set => _routableCidrs = value;
        }

        /// <summary>
        /// The ID of the Heroku Private Space where the VPN connection will be established.
        /// </summary>
        [Input("space")]
        public Input<string>? Space { get; set; }

        /// <summary>
        /// The CIDR block for the Heroku Private Space. The network where the VPN will be established should be configured to route traffic destined for this CIDR block over the VPN link.
        /// </summary>
        [Input("spaceCidrBlock")]
        public Input<string>? SpaceCidrBlock { get; set; }

        [Input("tunnels")]
        private InputList<Inputs.VpnConnectionTunnelGetArgs>? _tunnels;

        /// <summary>
        /// Details about each VPN tunnel endpoint.
        /// </summary>
        public InputList<Inputs.VpnConnectionTunnelGetArgs> Tunnels
        {
            get => _tunnels ?? (_tunnels = new InputList<Inputs.VpnConnectionTunnelGetArgs>());
            set => _tunnels = value;
        }

        public VpnConnectionState()
        {
        }
        public static new VpnConnectionState Empty => new VpnConnectionState();
    }
}
